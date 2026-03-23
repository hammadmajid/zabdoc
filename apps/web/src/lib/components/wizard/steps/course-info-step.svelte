<script lang="ts">
	import { formStore } from "$lib/stores/form-store.svelte";
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import * as Select from "$lib/components/ui/select/index.js";
	import Input from "$lib/components/ui/input/input.svelte";
	import Label from "$lib/components/ui/label/label.svelte";
	import CourseSelector from "./course-selector.svelte";
	import BookOpen from "@lucide/svelte/icons/book-open";
	import GraduationCap from "@lucide/svelte/icons/graduation-cap";
	import ClipboardList from "@lucide/svelte/icons/clipboard-list";
	import { scale } from "svelte/transition";
	import { quintOut } from "svelte/easing";

	// Initialize from localStorage on mount
	$effect(() => {
		formStore.initFromLocalStorage();
	});

	// Reset selectedCourse if it's no longer in the filtered courses
	// This happens when the document type changes (Assignment vs Lab Task)
	$effect(() => {
		// Access both courses and wizardStore.documentType to trigger reactivity
		formStore.courses;
		wizardStore.documentType;
		formStore.validateSelectedCourse();
	});
</script>

<div class="mx-auto flex w-full max-w-2xl flex-col items-center px-4 py-8">
	<div class="neo-border neo-shadow-lg mb-8 -rotate-1 bg-secondary px-6 py-3">
		<div class="flex items-center gap-3">
			<BookOpen class="size-6" />
			<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">Course Details</h1>
		</div>
	</div>

	<p class="mb-8 max-w-md text-center text-base font-medium">
		Select your class and course. Instructor details will be filled automatically.
	</p>

	<div class="w-full space-y-6">
		<!-- Class Selection -->
		<div
			class="neo-border neo-shadow space-y-4 bg-card p-6"
			in:scale={{ duration: 300, delay: 100, easing: quintOut }}
		>
			<div class="mb-2 flex items-center gap-2">
				<div class="neo-border-sm bg-primary p-2">
					<GraduationCap class="size-5" />
				</div>
				<span class="text-sm font-bold uppercase">Select Class</span>
			</div>

			<Select.Root type="single" name="class" required bind:value={formStore.selectedClass}>
				<Select.Trigger
					class="neo-border-sm h-auto min-h-12 w-full! text-lg whitespace-normal"
				>
					<span class="text-left leading-tight">{formStore.classTriggerContent}</span>
				</Select.Trigger>
				<Select.Content>
					<Select.Group>
						<Select.Label>Available Classes</Select.Label>
						{#each formStore.classes as classItem (classItem.value)}
							<Select.Item value={classItem.value} label={classItem.label}>
								{classItem.label}
							</Select.Item>
						{/each}
					</Select.Group>
				</Select.Content>
			</Select.Root>
		</div>

		<!-- Course Selection -->
		<div
			class="neo-border neo-shadow space-y-4 bg-card p-6"
			in:scale={{ duration: 300, delay: 200, easing: quintOut }}
		>
			<div class="mb-2 flex items-center gap-2">
				<div class="neo-border-sm bg-accent p-2">
					<BookOpen class="size-5" />
				</div>
				<span class="text-sm font-bold uppercase">Select Course</span>
			</div>

			<CourseSelector
				bind:value={formStore.selectedCourse}
				courses={formStore.courses}
				disabled={!formStore.selectedClass}
			/>

			{#if !formStore.selectedClass}
				<p class="neo-border-sm bg-muted px-3 py-2 text-sm text-muted-foreground">
					Please select a class first to see available courses.
				</p>
			{/if}
		</div>

		<!-- Auto-filled Course Details -->
		{#if formStore.courseDetails}
			<div
				class="neo-border neo-shadow space-y-4 bg-muted p-6"
				in:scale={{ duration: 300, easing: quintOut }}
			>
				<div class="mb-2 flex items-center gap-2">
					<div class="neo-border-sm bg-secondary p-2">
						<ClipboardList class="size-5" />
					</div>
					<span class="text-sm font-bold uppercase">Auto-filled Details</span>
				</div>

				<div class="grid gap-4 sm:grid-cols-2">
					<div class="space-y-2">
						<Label for="instructor" class="text-xs font-bold uppercase"
							>Instructor</Label
						>
						<Input
							id="instructor"
							name="instructor"
							type="text"
							value={formStore.courseDetails.instructor}
							readonly
							class="neo-border-sm bg-card"
						/>
					</div>

					<div class="space-y-2">
						<Label for="course-code" class="text-xs font-bold uppercase"
							>Course Code</Label
						>
						<Input
							id="course-code"
							name="courseCode"
							type="text"
							value={formStore.courseDetails.code}
							readonly
							class="neo-border-sm bg-card"
						/>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>
