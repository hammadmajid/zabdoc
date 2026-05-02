<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import DocumentTypeStep from "./steps/document-type-step.svelte";
	import TeamTypeStep from "./steps/team-type-step.svelte";
	import StudentInfoStep from "./steps/student-info-step.svelte";
	import CourseInfoStep from "./steps/course-info-step.svelte";
	import DocumentInfoStep from "./steps/document-info-step.svelte";
	import FinalizeStep from "./steps/finalize-step.svelte";
	import Button from "$lib/components/ui/button/button.svelte";
	import ArrowLeft from "@lucide/svelte/icons/arrow-left";
	import ArrowRight from "@lucide/svelte/icons/arrow-right";
	import RotateCcw from "@lucide/svelte/icons/rotate-ccw";
	import { fly, fade } from "svelte/transition";
	import { cubicOut } from "svelte/easing";

	let { baseURL }: { baseURL: string } = $props();
</script>

<div class="flex min-h-[60vh] flex-col">
	<!-- Progress Bar -->
	{#if wizardStore.currentStep !== "select-document"}
		<div class="mb-8" in:fade={{ duration: 300 }}>
			<div class="mb-2 flex items-center justify-between">
				<span class="text-sm font-bold uppercase">
					Step {wizardStore.currentStepIndex + 1} of {wizardStore.totalSteps}
				</span>
				<button
					type="button"
					onclick={() => wizardStore.reset()}
					class="flex cursor-pointer items-center gap-1 text-sm font-bold text-muted-foreground uppercase transition-colors hover:text-foreground"
				>
					<RotateCcw class="size-4" />
					Start Over
				</button>
			</div>
			<div class="neo-border-sm h-4 overflow-hidden bg-muted">
				<div
					class="h-full bg-primary transition-all duration-500 ease-out"
					style="width: {wizardStore.progress}%"
				></div>
			</div>
			<div class="mt-2 flex justify-between">
				{#each wizardStore.stepOrder as step, i}
					<button
						type="button"
						onclick={() => {
							if (i < wizardStore.currentStepIndex) {
								wizardStore.goToStep(step);
							}
						}}
						disabled={i > wizardStore.currentStepIndex}
						class="neo-border-sm h-6 w-6 cursor-pointer text-xs font-bold transition-all disabled:cursor-not-allowed"
						class:bg-primary={i <= wizardStore.currentStepIndex}
						class:bg-muted={i > wizardStore.currentStepIndex}
						class:hover:bg-secondary={i < wizardStore.currentStepIndex}
					>
						{i + 1}
					</button>
				{/each}
			</div>
		</div>
	{/if}

	<!-- Step Content -->
	<div class="relative flex-1 overflow-hidden">
		{#key wizardStore.currentStep}
			<div
				class="w-full"
				in:fly={{
					x: wizardStore.direction === "forward" ? 40 : -40,
					duration: 250,
					easing: cubicOut,
					delay: 150
				}}
				out:fly={{
					x: wizardStore.direction === "forward" ? -40 : 40,
					duration: 150,
					easing: cubicOut
				}}
			>
				{#if wizardStore.currentStep === "select-document"}
					<DocumentTypeStep />
				{:else if wizardStore.currentStep === "select-team"}
					<TeamTypeStep />
				{:else if wizardStore.currentStep === "student-info"}
					<StudentInfoStep />
				{:else if wizardStore.currentStep === "course-info"}
					<CourseInfoStep />
				{:else if wizardStore.currentStep === "document-info"}
					<DocumentInfoStep />
				{:else if wizardStore.currentStep === "finalize"}
					<FinalizeStep {baseURL} />
				{/if}
			</div>
		{/key}
	</div>

	<!-- Navigation Buttons -->
	{#if wizardStore.currentStep !== "select-document" && wizardStore.currentStep !== "finalize"}
		<div class="mt-auto flex items-center justify-between pt-8" in:fade={{ duration: 300 }}>
			<button
				type="button"
				onclick={() => wizardStore.prevStep()}
				disabled={wizardStore.isFirstStep}
				class="neo-border neo-shadow flex cursor-pointer items-center gap-2 bg-card px-6 py-3 font-bold uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none disabled:cursor-not-allowed disabled:opacity-50 disabled:hover:translate-x-0 disabled:hover:translate-y-0"
			>
				<ArrowLeft class="size-5" />
				Back
			</button>

			<button
				type="button"
				onclick={() => wizardStore.nextStep()}
				disabled={!wizardStore.canProceed()}
				class="neo-border neo-shadow flex cursor-pointer items-center gap-2 bg-primary px-6 py-3 font-bold text-primary-foreground uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none disabled:cursor-not-allowed disabled:opacity-50 disabled:hover:translate-x-0 disabled:hover:translate-y-0"
			>
				Next
				<ArrowRight class="size-5" />
			</button>
		</div>
	{/if}

	<!-- Back button only for finalize -->
	{#if wizardStore.currentStep === "finalize"}
		<div class="mt-auto flex items-center justify-start pt-8" in:fade={{ duration: 300 }}>
			<button
				type="button"
				onclick={() => wizardStore.prevStep()}
				class="neo-border neo-shadow flex cursor-pointer items-center gap-2 bg-card px-6 py-3 font-bold uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
			>
				<ArrowLeft class="size-5" />
				Back
			</button>
		</div>
	{/if}
</div>
