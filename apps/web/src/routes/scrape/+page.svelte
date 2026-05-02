<script lang="ts">
	import { fade, scale } from "svelte/transition";
	import { quintOut } from "svelte/easing";
	import ArrowLeft from "@lucide/svelte/icons/arrow-left";
	import Database from "@lucide/svelte/icons/database";
	import Loader2 from "@lucide/svelte/icons/loader-2";
	import XCircle from "@lucide/svelte/icons/x-circle";
	import * as Card from "$lib/components/ui/card";
	import * as Tabs from "$lib/components/ui/tabs";
	import * as Table from "$lib/components/ui/table";
	import * as Select from "$lib/components/ui/select";
	import Button from "$lib/components/ui/button/button.svelte";
	import Checkbox from "$lib/components/ui/checkbox/checkbox.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import Label from "$lib/components/ui/label/label.svelte";
	import loadingMessages from "$lib/loading-msgs";
	import type { PageProps } from "./$types";
	import type { CourseData, MarkCategory, MarkEntry, MarkGroup } from "$lib/types";
	import { scrapeInitSchema, scrapedData } from "@repo/types";

	function getDefaultSemester(): "Fall" | "Spring" | "Summer" {
		const month = new Date().getMonth() + 1; // 1-12
		if (month >= 9 || month <= 1) {
			return "Fall"; // Sep-Jan: 9, 10, 11, 12, 1
		} else if (month >= 2 && month <= 6) {
			return "Spring"; // Feb-June: 2, 3, 4, 5, 6
		} else {
			return "Summer"; // July-Aug: 7, 8
		}
	}

	let username = $state("");
	let password = $state("");
	let campus = $state("Islamabad");
	let semester = $state<"Fall" | "Spring" | "Summer">(getDefaultSemester());
	let isLoading = $state(false);
	let isError = $state(false);
	let errorMessage = $state("");
	let loadingMessage = $state("");
	let loadingInterval: ReturnType<typeof setInterval> | null = null;
	let courseData = $state<CourseData[]>([]);
	let hasResults = $state(false);
	let agreedToTerms = $state(false);
	let { data }: PageProps = $props();

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

	/**
	 * Convert display values to schema-compliant values
	 */
	function convertFormValuesToSchema(displaySemester: string): 'fall' | 'spring' | 'summer' {
		const semesterMap: Record<string, 'fall' | 'spring' | 'summer'> = {
			'Fall': 'fall',
			'Spring': 'spring',
			'Summer': 'summer'
		};
		return semesterMap[displaySemester] as 'fall' | 'spring' | 'summer';
	}

	async function handleScrap() {
		if (!username.trim() || !password.trim()) {
			isError = true;
			errorMessage = "Username and password are required";
			return;
		}

		if (!agreedToTerms) {
			isError = true;
			errorMessage = "You must agree to the Terms of Use and Privacy Policy to continue.";
			return;
		}

		isLoading = true;
		isError = false;
		errorMessage = "";
		startLoadingMessages();

		try {
			const apiUrl = `${data.baseURL}/scrape`;

			// Validate request payload against scrapeInitSchema
			const requestPayload = {
				username,
				password,
				campus: 'isb' as const, // Match schema enum
				semester: convertFormValuesToSchema(semester),
			};

			const schemaValidation = scrapeInitSchema.safeParse(requestPayload);
			if (!schemaValidation.success) {
				isError = true;
				errorMessage = `Invalid request data: ${schemaValidation.error.message}`;
				return;
			}

			const response = await fetch(apiUrl, {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(requestPayload)
			});

			if (!response.ok) {
				if (response.status === 400) {
					isError = true;
					errorMessage = "Invalid request data. Please check your input.";
					return;
				} else if (response.status === 500) {
					isError = true;
					errorMessage = "Server error. Failed to fetch attendance and marks data.";
					return;
				} else {
					isError = true;
					errorMessage = `Error: ${response.status}`;
					return;
				}
			}

			const result = await response.json();

			console.log("API Response:", result);

			// Validate response against scrapedData schema
			const validatedData = scrapedData.safeParse(result);
			if (!validatedData.success) {
				isError = true;
				errorMessage = `Invalid response data from server: ${validatedData.error.message}`;
				console.error("Validation errors:", validatedData.error);
				return;
			}

			// Transform validated data into CourseData array for display
			courseData = Object.entries(validatedData.data).map(
				([courseName, courseValue]) => {
					const attendance = courseValue.attendance;
					const marks = courseValue.marks;

					return {
						courseName,
						instructor: attendance.instructor || "Unknown Instructor",
						records: attendance.records || [],
						marks: marks.entries || []
					};
				}
			);

			console.log("Processed courseData:", courseData);
			hasResults = true;
		} catch (error) {
			isError = true;
			errorMessage = error instanceof Error ? error.message : String(error);
			console.error("Scrape error:", error);
		} finally {
			stopLoadingMessages();
			loadingMessage = "";
			isLoading = false;
		}
	}

	function getStatusColor(status: string): string {
		const normalizedStatus = status.toLowerCase().trim();
		if (normalizedStatus === "present" || normalizedStatus === "p") {
			return "text-green-600 bg-green-100 dark:bg-green-950";
		} else if (normalizedStatus === "absent" || normalizedStatus === "a") {
			return "text-red-600 bg-red-100 dark:bg-red-950";
		} else if (normalizedStatus === "late" || normalizedStatus === "l") {
			return "text-yellow-600 bg-yellow-100 dark:bg-yellow-950";
		}
		return "text-muted-foreground bg-muted";
	}

	function getMarkCellClass(obtained: string): string {
		if (obtained.toLowerCase() === "not entered") {
			return "text-muted-foreground italic";
		}
		return "";
	}

	function parseMarkValue(value: string): number | null {
		const normalized = value.trim().toLowerCase();
		if (!normalized || normalized === "not entered") {
			return null;
		}

		const parsed = Number.parseFloat(value);
		return Number.isFinite(parsed) ? parsed : null;
	}

	function formatMarkValue(value: number): string {
		if (Number.isInteger(value)) {
			return `${value}`;
		}

		return value
			.toFixed(2)
			.replace(/\.00$/, "")
			.replace(/(\.\d)0$/, "$1");
	}

	function getMarkCategory(head: string): MarkCategory {
		const normalized = head.trim().toLowerCase();
		if (normalized.includes("quiz")) return "quiz";
		if (normalized.includes("assignment")) return "assignment";
		if (normalized.includes("lab") || normalized.includes("practical")) return "lab";
		if (normalized.includes("project")) return "project";
		if (
			normalized.includes("mid") ||
			normalized.includes("final") ||
			normalized.includes("paper") ||
			normalized.includes("exam")
		) {
			return "exam";
		}
		return "other";
	}

	function groupMarks(marks: MarkEntry[]): {
		groups: MarkGroup[];
		overallMax: number;
		overallObtained: number | null;
	} {
		const categoryOrder: MarkCategory[] = [
			"quiz",
			"assignment",
			"lab",
			"project",
			"exam",
			"other"
		];
		const categoryLabels: Record<MarkCategory, string> = {
			quiz: "Quizzes",
			assignment: "Assignments",
			lab: "Lab Work",
			project: "Projects",
			exam: "Exams",
			other: "Other"
		};

		const groupMap = new Map<
			MarkCategory,
			{ entries: MarkEntry[]; totalMax: number; totalObtained: number; hasObtained: boolean }
		>();

		let overallMax = 0;
		let overallObtained = 0;
		let hasOverallObtained = false;

		for (const mark of marks) {
			const category = getMarkCategory(mark.head);
			if (!groupMap.has(category)) {
				groupMap.set(category, {
					entries: [],
					totalMax: 0,
					totalObtained: 0,
					hasObtained: false
				});
			}
			const group = groupMap.get(category)!;

			const max = parseMarkValue(mark.max) ?? 0;
			const obtained = parseMarkValue(mark.obtained);

			group.entries.push(mark);
			group.totalMax += max;
			if (obtained !== null) {
				group.totalObtained += obtained;
				group.hasObtained = true;
			}

			overallMax += max;
			if (obtained !== null) {
				overallObtained += obtained;
				hasOverallObtained = true;
			}
		}

		const groups: MarkGroup[] = [];
		for (const category of categoryOrder) {
			const data = groupMap.get(category);
			if (!data) continue;
			groups.push({
				category,
				label: categoryLabels[category],
				entries: data.entries,
				totalMax: data.totalMax,
				totalObtained: data.hasObtained ? data.totalObtained : null
			});
		}

		return {
			groups,
			overallMax,
			overallObtained: hasOverallObtained ? overallObtained : null
		};
	}

	// Disable this feature
	let disabled = $state(true);

</script>

<div class="flex min-h-[70vh] flex-col">
	<!-- Back to Home Button -->
	<div class="mb-4" in:fade={{ duration: 300 }}>
		<a
			href="/"
			class="neo-border neo-shadow inline-flex w-fit items-center gap-2 bg-card px-4 py-2 text-sm font-bold uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
		>
			<ArrowLeft class="size-4" />
			Back to Home
		</a>
	</div>

	{#if disabled}
		<div class="mx-auto flex w-full max-w-lg flex-col items-center justify-center px-4 py-16">
			<Card.Root class="w-full border-dashed">
				<Card.Content class="flex flex-col items-center text-center py-12">
					<div class="mb-4 rounded-full bg-muted p-4">
						<XCircle class="size-8 text-muted-foreground" />
					</div>
					<Card.Title class="mb-2 text-2xl">Feature Not Available</Card.Title>
					<Card.Description class="text-base">
						This feature is currently under maintenance. Please check back soon.
					</Card.Description>
				</Card.Content>
			</Card.Root>
		</div>
	{:else}
		{#if !hasResults}
		<!-- Login Form -->
		<div
			class="mx-auto flex w-full max-w-md flex-1 flex-col items-center justify-center px-4 py-8"
		>
			<div class="neo-border neo-shadow-lg mb-8 -rotate-2 bg-primary px-8 py-4">
				<div class="flex items-center gap-3">
					<Database class="size-6" />
					<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
						Scrape Data
					</h1>
				</div>
			</div>

			<p class="mb-8 max-w-md text-center text-base font-medium">
				Enter your ZabDesk credentials to fetch attendance and marks data.
			</p>

			{#if isError}
				<div
					class="neo-border neo-shadow-lg mb-6 w-full border-destructive bg-destructive/10 p-4"
					in:scale={{ duration: 400, easing: quintOut }}
				>
					<div class="flex items-center gap-3">
						<XCircle class="size-5 shrink-0 text-destructive" />
						<p class="text-sm font-medium text-destructive">
							{errorMessage}
						</p>
					</div>
				</div>
			{/if}

			<div class="w-full space-y-4">
				<div class="space-y-2">
					<Label for="username" class="text-sm font-bold uppercase">Username</Label>
					<Input
						id="username"
						type="text"
						placeholder="Enter your username"
						bind:value={username}
						disabled={isLoading}
						class="neo-border-sm"
					/>
				</div>

				<div class="space-y-2">
					<Label for="password" class="text-sm font-bold uppercase">Password</Label>
					<Input
						id="password"
						type="password"
						placeholder="Enter your password"
						bind:value={password}
						disabled={isLoading}
						class="neo-border-sm"
						onkeydown={(e) => {
							if (e.key === "Enter") {
								handleScrap();
							}
						}}
					/>
				</div>

				<div class="flex gap-2">
					<div class="w-full space-y-2">
						<Label for="campus" class="text-sm font-bold uppercase">Campus</Label>
						<Select.Root type="single" bind:value={campus} disabled={isLoading}>
							<Select.Trigger class="neo-border-sm w-full">
								{campus}
							</Select.Trigger>
							<Select.Content>
								<Select.Item value="Islamabad">Islamabad</Select.Item>
							</Select.Content>
						</Select.Root>
					</div>

					<div class="w-full space-y-2">
						<Label for="semester" class="text-sm font-bold uppercase">Semester</Label>
						<Select.Root type="single" bind:value={semester} disabled={isLoading}>
							<Select.Trigger class="neo-border-sm w-full">
								{semester}
							</Select.Trigger>
							<Select.Content>
								<Select.Item value="Fall">Fall</Select.Item>
								<Select.Item value="Spring">Spring</Select.Item>
								<Select.Item value="Summer">Summer</Select.Item>
							</Select.Content>
						</Select.Root>
					</div>
				</div>

				<!-- Terms & Privacy Agreement -->
				<div class="neo-border neo-shadow-sm my-6 flex items-start gap-3 bg-card p-4">
					<Checkbox
						id="scrape-terms-agree"
						bind:checked={agreedToTerms}
						disabled={isLoading}
						class="mt-1"
					/>
					<label
						for="scrape-terms-agree"
						class="flex-1 cursor-pointer text-sm font-medium"
					>
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

				<Button
					onclick={handleScrap}
					disabled={isLoading}
					class="neo-border neo-shadow mt-6 w-full bg-primary px-6 py-6 text-lg font-black text-primary-foreground uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					{#if isLoading}
						<Loader2 class="mr-2 size-5 animate-spin" />
						<span class="text-sm font-medium tracking-normal normal-case">
							{loadingMessage || "Scraping Data..."}
						</span>
					{:else}
						<Database class="mr-2 size-5" />
						Scrape Data
					{/if}
				</Button>
			</div>
		</div>
	{:else}
		<!-- Results Display -->
		<div class="mx-auto flex w-full max-w-6xl flex-col px-4 py-8">
			<div class="mb-8 flex items-center justify-between">
				<div class="neo-border neo-shadow-lg -rotate-1 bg-primary px-6 py-3">
					<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
						Scrapped Data
					</h1>
				</div>
			</div>

			{#if courseData.length === 0}
				<div class="py-12 text-center">
					<p class="text-lg font-medium text-muted-foreground">
						No data found.
					</p>
				</div>
			{:else}
				<div class="grid gap-6 md:grid-cols-2">
					{#each courseData as course, index}
						<div in:scale={{ duration: 300, delay: index * 50, easing: quintOut }}>
							<Card.Root>
								<Card.Header>
									<Card.Title class="text-xl font-black uppercase">
										{course.courseName}
									</Card.Title>
									<Card.Description class="font-medium">
										{course.instructor}
									</Card.Description>
								</Card.Header>
								<Card.Content>
									<Tabs.Root value="attendance" class="w-full">
										<Tabs.List class="neo-border-sm grid w-full grid-cols-2">
											<Tabs.Trigger
												value="attendance"
												class="text-sm font-bold uppercase"
											>
												Attendance
											</Tabs.Trigger>
											<Tabs.Trigger
												value="marks"
												class="text-sm font-bold uppercase"
											>
												Marks
											</Tabs.Trigger>
										</Tabs.List>
										<Tabs.Content value="attendance" class="mt-4">
											{#if !course.records || course.records.length === 0}
												<p
													class="py-4 text-center text-sm text-muted-foreground"
												>
													No attendance records available.
												</p>
											{:else}
												<Table.Root>
													<Table.Header>
														<Table.Row>
															<Table.Head>Lecture</Table.Head>
															<Table.Head>Date</Table.Head>
															<Table.Head>Status</Table.Head>
														</Table.Row>
													</Table.Header>
													<Table.Body>
														{#each course.records as record}
															<Table.Row>
																<Table.Cell class="font-medium">
																	{record.lecture}
																</Table.Cell>
																<Table.Cell
																	class="text-muted-foreground"
																>
																	{record.date}
																</Table.Cell>
																<Table.Cell>
																	<span
																		class="inline-block rounded px-2 py-1 text-xs font-bold uppercase {getStatusColor(
																			record.status
																		)}"
																	>
																		{record.status}
																	</span>
																</Table.Cell>
															</Table.Row>
														{/each}
													</Table.Body>
												</Table.Root>
											{/if}
										</Tabs.Content>
										<Tabs.Content value="marks" class="mt-4">
											{#if !course.marks || course.marks.length === 0}
												<p
													class="py-4 text-center text-sm text-muted-foreground"
												>
													No marks records available.
												</p>
											{:else}
												{@const { groups, overallMax, overallObtained } =
													groupMarks(course.marks)}
												<Table.Root>
													<Table.Header>
														<Table.Row>
															<Table.Head>Marks Head</Table.Head>
															<Table.Head>Max</Table.Head>
															<Table.Head>Obtained</Table.Head>
														</Table.Row>
													</Table.Header>
													<Table.Body>
														{#each groups as group}
															<!-- Category section header -->
															<Table.Row
																class="bg-primary/10 hover:bg-primary/10"
															>
																<Table.Cell
																	colspan={3}
																	class="py-2 text-xs font-black tracking-widest text-primary uppercase"
																>
																	{group.label}
																</Table.Cell>
															</Table.Row>
															<!-- Individual entries -->
															{#each group.entries as mark}
																<Table.Row>
																	<Table.Cell
																		class="pl-6 font-medium"
																	>
																		{mark.head}
																	</Table.Cell>
																	<Table.Cell
																		class="text-muted-foreground"
																	>
																		{mark.max}
																	</Table.Cell>
																	<Table.Cell
																		class={getMarkCellClass(
																			mark.obtained
																		)}
																	>
																		{mark.obtained}
																	</Table.Cell>
																</Table.Row>
															{/each}
															<!-- Category total -->
															<Table.Row class="border-t bg-muted/60">
																<Table.Cell
																	class="pl-6 text-xs font-bold tracking-wide text-muted-foreground uppercase"
																>
																	{group.label} Total
																</Table.Cell>
																<Table.Cell class="font-bold">
																	{formatMarkValue(
																		group.totalMax
																	)}
																</Table.Cell>
																<Table.Cell
																	class={group.totalObtained ===
																	null
																		? "font-bold text-muted-foreground italic"
																		: "font-bold"}
																>
																	{group.totalObtained === null
																		? "Not Entered"
																		: formatMarkValue(
																				group.totalObtained
																			)}
																</Table.Cell>
															</Table.Row>
														{/each}
														<!-- Overall total -->
														<Table.Row
															class="border-t-2 bg-primary/15 hover:bg-primary/15"
														>
															<Table.Cell
																class="text-xs font-black tracking-widest text-primary uppercase"
															>
																Overall Total
															</Table.Cell>
															<Table.Cell
																class="font-black text-primary"
															>
																{formatMarkValue(overallMax)}
															</Table.Cell>
															<Table.Cell
																class={overallObtained === null
																	? "font-black text-muted-foreground italic"
																	: "font-black text-primary"}
															>
																{overallObtained === null
																	? "Not Entered"
																	: formatMarkValue(
																			overallObtained
																		)}
															</Table.Cell>
														</Table.Row>
													</Table.Body>
												</Table.Root>
											{/if}
										</Tabs.Content>
									</Tabs.Root>
								</Card.Content>
							</Card.Root>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
{/if}
</div>
