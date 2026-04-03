<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import { formStore } from "$lib/stores/form-store.svelte";
	import { smartName } from "$lib/smart-name";
	import loadingMessages from "$lib/loading-msgs";
	import Button from "$lib/components/ui/button/button.svelte";
	import Checkbox from "$lib/components/ui/checkbox/checkbox.svelte";
	import Loader2 from "@lucide/svelte/icons/loader-2";
	import CheckCircle from "@lucide/svelte/icons/check-circle";
	import XCircle from "@lucide/svelte/icons/x-circle";
	import FileQuestion from "@lucide/svelte/icons/file-question";
	import { scale } from "svelte/transition";
	import { quintOut } from "svelte/easing";
	import FileText from "@lucide/svelte/icons/file-text";
	import Download from "@lucide/svelte/icons/download";
	import Pencil from "@lucide/svelte/icons/pencil";
	import Users from "@lucide/svelte/icons/users";
	import User from "@lucide/svelte/icons/user";
	import BookOpen from "@lucide/svelte/icons/book-open";
	import CalendarDays from "@lucide/svelte/icons/calendar-days";
	import RefreshCw from "@lucide/svelte/icons/refresh-cw";

	let isLoading = $state(false);
	let isSuccess = $state(false);
	let isError = $state(false);
	let errorMessage = $state("");
	let termsErrorMessage = $state("");
	let loadingMessage = $state("");
	let loadingInterval: ReturnType<typeof setInterval> | null = null;
	let agreedToTerms = $state(false);
	let { baseURL }: { baseURL: string } = $props();

	function getRandomLoadingMessage() {
		return loadingMessages[Math.floor(Math.random() * loadingMessages.length)];
	}

	function startLoadingMessages() {
		loadingMessage = getRandomLoadingMessage();
		loadingInterval = setInterval(() => {
			loadingMessage = getRandomLoadingMessage();
		}, 2000);
	}

	function stopLoadingMessages() {
		if (loadingInterval) {
			clearInterval(loadingInterval);
			loadingInterval = null;
		}
	}

	async function handleSubmit() {
		termsErrorMessage = "";
		if (!agreedToTerms) {
			isError = false;
			errorMessage = "";
			termsErrorMessage = "You must agree to the Terms of Use and Privacy Policy to continue.";
			return;
		}

		isLoading = true;
		isSuccess = false;
		isError = false;
		errorMessage = "";
		startLoadingMessages();

		try {
			const jsonData = formStore.buildJSON();

			const apiUrl = `${baseURL}/document`;

			// Create AbortController for timeout
			const controller = new AbortController();
			const timeoutId = setTimeout(() => controller.abort(), 30000); // 30 second timeout

			try {
				const response = await fetch(apiUrl, {
					method: "POST",
					headers: {
						"Content-Type": "application/json"
					},
					body: JSON.stringify(jsonData),
					signal: controller.signal
				});

				clearTimeout(timeoutId);

				if (!response.ok) {
					let errorText = "";
					try {
						errorText = await response.text();
					} catch {
						// If we can't read the response body, use a default message
						errorText = "";
					}

					// Provide more specific error messages based on status code
					if (response.status === 400) {
						isError = true;
						errorMessage =
							"Invalid request. Please check your data and try again.";
						return;
					} else if (response.status === 500) {
						isError = true;
						errorMessage = "Server error. Please try again later.";
						return;
					} else {
						isError = true;
						errorMessage = errorText || `Server error: ${response.status}`;
						return;
					}
				}

				const blob = await response.blob();
				const url = window.URL.createObjectURL(blob);
				const a = document.createElement("a");
				a.style.display = "none";
				a.href = url;
				a.download = smartName(jsonData);
				document.body.appendChild(a);
				a.click();
				window.URL.revokeObjectURL(url);
				document.body.removeChild(a);

				isSuccess = true;
			} catch (fetchError) {
				clearTimeout(timeoutId);

				// Check if it was a timeout/abort error
				if (fetchError instanceof Error && fetchError.name === "AbortError") {
					isError = true;
					errorMessage =
						"Request timed out. The server took too long to respond. Please try again.";
					return;
				}

				isError = true;
				errorMessage = "Something went wrong";
			}
		} catch (error) {
			isError = true;
			errorMessage = error instanceof Error ? error.message : String(error);
		} finally {
			stopLoadingMessages();
			isLoading = false;
		}
	}

	function handleStartOver() {
		isSuccess = false;
		isError = false;
		errorMessage = "";
		termsErrorMessage = "";
		wizardStore.reset();
	}

	function handleRetry() {
		isError = false;
		errorMessage = "";
		termsErrorMessage = "";
		handleSubmit();
	}
</script>

<div class="mx-auto flex w-full max-w-3xl flex-col items-center px-4 py-8">
	<div class="neo-border neo-shadow-lg mb-8 -rotate-1 bg-primary px-6 py-3">
		<div class="flex items-center gap-3">
			<CheckCircle class="size-6" />
			<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
				Review & Generate
			</h1>
		</div>
	</div>

	<p class="mb-8 max-w-md text-center text-base font-medium">
		Make sure everything looks correct before generating your PDF.
	</p>

	{#if isSuccess}
		<div
			class="neo-border neo-shadow-lg w-full space-y-6 bg-secondary p-8 text-center"
			in:scale={{ duration: 400, easing: quintOut }}
		>
			<div class="neo-border-sm mx-auto inline-block bg-card p-4">
				<CheckCircle class="size-16 text-green-600" />
			</div>
			<h2 class="text-2xl font-black uppercase">Success!</h2>
			<p class="text-muted-foreground">Your PDF has been generated and downloaded.</p>
			<div class="flex flex-col justify-center gap-4 pt-4 sm:flex-row">
				<Button
					onclick={handleSubmit}
					class="neo-border neo-shadow bg-primary transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					<Download class="mr-2 size-5" />
					Download Again
				</Button>
				<Button
					variant="outline"
					onclick={handleStartOver}
					class="neo-border neo-shadow transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					Start Over
				</Button>
			</div>
		</div>
	{:else if isError}
		<div
			class="neo-border neo-shadow-lg w-full space-y-6 bg-destructive/10 p-8 text-center"
			in:scale={{ duration: 400, easing: quintOut }}
		>
			<div class="neo-border-sm mx-auto inline-block border-destructive bg-card p-4">
				<XCircle class="size-16 text-destructive" />
			</div>
			<h2 class="text-2xl font-black text-destructive uppercase">Error!</h2>
			<div class="space-y-2">
				<p class="text-muted-foreground">Something went wrong while generating your PDF.</p>
				<div class="neo-border-sm bg-card p-3 text-left">
					<p class="font-mono text-sm break-all text-destructive">
						{errorMessage}
					</p>
				</div>
			</div>
			<div class="flex flex-col justify-center gap-4 pt-4 sm:flex-row">
				<Button
					onclick={handleRetry}
					class="neo-border neo-shadow bg-primary transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					<RefreshCw class="mr-2 size-5" />
					Try Again
				</Button>
				<Button
					variant="outline"
					onclick={handleStartOver}
					class="neo-border neo-shadow transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					Start Over
				</Button>
			</div>
		</div>
	{:else}
		<div class="w-full space-y-4">
			<!-- Document Type -->
			<div
				class="neo-border neo-shadow flex items-center justify-between bg-card p-4"
				in:scale={{ duration: 300, delay: 50, easing: quintOut }}
			>
				<div class="flex items-center gap-3">
					<div class="neo-border-sm bg-primary p-2">
						<FileText class="size-5" />
					</div>
					<div>
						<p class="text-xs font-bold text-muted-foreground uppercase">
							Document Type
						</p>
						<p class="font-bold">{wizardStore.documentType}</p>
					</div>
				</div>
				<button
					type="button"
					onclick={() => wizardStore.goToStep("select-document")}
					class="neo-border-sm cursor-pointer bg-muted p-2 transition-colors hover:bg-secondary"
				>
					<Pencil class="size-4" />
				</button>
			</div>

			<!-- Student Info -->
			<div
				class="neo-border neo-shadow flex items-center justify-between bg-card p-4"
				in:scale={{ duration: 300, delay: 100, easing: quintOut }}
			>
				<div class="flex items-center gap-3">
					<div class="neo-border-sm bg-secondary p-2">
						{#if wizardStore.teamType === "blank"}
							<FileQuestion class="size-5" />
						{:else if wizardStore.teamType === "individual"}
							<User class="size-5" />
						{:else}
							<Users class="size-5" />
						{/if}
					</div>
					<div>
						<p class="text-xs font-bold text-muted-foreground uppercase">
							{#if wizardStore.teamType === "blank"}
								Student Info
							{:else if wizardStore.teamType === "individual"}
								Student
							{:else}
								Group Members
							{/if}
						</p>
						{#if wizardStore.teamType === "blank"}
							<p class="font-bold">Left blank</p>
							<p class="text-sm text-muted-foreground">Fill in after printing</p>
						{:else if wizardStore.teamType === "individual"}
							<p class="font-bold">{formStore.studentName || "Not set"}</p>
							<p class="text-sm text-muted-foreground">
								{formStore.regNo || "No reg number"}
							</p>
						{:else}
							<p class="font-bold">{formStore.students.length} members</p>
							<p class="text-sm text-muted-foreground">
								{formStore.students.map((s) => s.name || "Unnamed").join(", ")}
							</p>
						{/if}
					</div>
				</div>
				<button
					type="button"
					onclick={() => wizardStore.goToStep("student-info")}
					class="neo-border-sm cursor-pointer bg-muted p-2 transition-colors hover:bg-secondary"
				>
					<Pencil class="size-4" />
				</button>
			</div>

			<!-- Course Info -->
			<div
				class="neo-border neo-shadow flex items-center justify-between bg-card p-4"
				in:scale={{ duration: 300, delay: 150, easing: quintOut }}
			>
				<div class="flex items-center gap-3">
					<div class="neo-border-sm bg-accent p-2">
						<BookOpen class="size-5" />
					</div>
					<div>
						<p class="text-xs font-bold text-muted-foreground uppercase">Course</p>
						<p class="font-bold">{formStore.selectedCourse || "Not selected"}</p>
						<p class="text-sm text-muted-foreground">
							{formStore.selectedClass || "No class"}
							• {formStore.courseDetails?.instructor || "No instructor"}
						</p>
					</div>
				</div>
				<button
					type="button"
					onclick={() => wizardStore.goToStep("course-info")}
					class="neo-border-sm cursor-pointer bg-muted p-2 transition-colors hover:bg-secondary"
				>
					<Pencil class="size-4" />
				</button>
			</div>

			<!-- Document Details -->
			<div
				class="neo-border neo-shadow flex items-center justify-between bg-card p-4"
				in:scale={{ duration: 300, delay: 200, easing: quintOut }}
			>
				<div class="flex items-center gap-3">
					<div class="neo-border-sm bg-primary p-2">
						<CalendarDays class="size-5" />
					</div>
					<div>
						<p class="text-xs font-bold text-muted-foreground uppercase">Details</p>
						<p class="font-bold">
							{wizardStore.documentType} #{formStore.document.number || "?"}
							• {formStore.document.marks || "?"} marks
						</p>
						<p class="text-sm text-muted-foreground">
							Due: {formStore.document.date || "Not set"}
						</p>
					</div>
				</div>
				<button
					type="button"
					onclick={() => wizardStore.goToStep("document-info")}
					class="neo-border-sm cursor-pointer bg-muted p-2 transition-colors hover:bg-secondary"
				>
					<Pencil class="size-4" />
				</button>
			</div>

			<!-- Images (Lab Task only) -->
		</div>

		<!-- Generate Button -->
		<div class="mt-8 w-full" in:scale={{ duration: 300, delay: 300, easing: quintOut }}>
			<!-- Terms & Privacy Agreement -->
			<div class="neo-border neo-shadow-sm mb-6 flex flex-col gap-2 bg-card p-4">
				<div class="flex items-start gap-3">
					<Checkbox
						id="finalize-terms-agree"
						bind:checked={agreedToTerms}
						disabled={isLoading}
						class="mt-1"
					/>
					<label for="finalize-terms-agree" class="flex-1 cursor-pointer text-sm font-medium">
						I agree to the <a
							href="/terms"
							target="_blank"
							rel="noreferrer"
							class="font-bold text-primary hover:underline">Terms of Use</a
						>
						and
						<a
							href="/privacy"
							target="_blank"
							rel="noreferrer"
							class="font-bold text-primary hover:underline">Privacy Policy</a
						>
					</label>
				</div>
				{#if termsErrorMessage}
					<p class="text-sm font-medium text-destructive">{termsErrorMessage}</p>
				{/if}
			</div>

			<button
				type="button"
				onclick={handleSubmit}
				disabled={isLoading}
				class="neo-border neo-shadow-lg flex w-full cursor-pointer items-center justify-center gap-3 bg-primary px-8 py-6 text-xl font-black tracking-wide text-primary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none disabled:cursor-not-allowed disabled:opacity-50 disabled:hover:translate-x-0 disabled:hover:translate-y-0 disabled:hover:shadow-(--neo-shadow-lg)"
			>
				{#if isLoading}
					<Loader2 class="size-6 animate-spin" />
					<span class="text-sm font-medium tracking-normal normal-case"
						>{loadingMessage}</span
					>
				{:else}
					<Download class="size-6" />
					Generate PDF
				{/if}
			</button>
		</div>
	{/if}
</div>
