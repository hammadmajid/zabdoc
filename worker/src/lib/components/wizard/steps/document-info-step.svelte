<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import { formStore } from "$lib/stores/form-store.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import DueDate from "$lib/components/forms/fields/due-date.svelte";
	import AssignmentNumberSelector from "./assignment-number-selector.svelte";
	import FileText from "@lucide/svelte/icons/file-text";
	import Hash from "@lucide/svelte/icons/hash";
	import Award from "@lucide/svelte/icons/award";
	import CalendarDays from "@lucide/svelte/icons/calendar-days";
	import { scale } from "svelte/transition";
	import { quintOut } from "svelte/easing";
</script>

<div class="mx-auto flex w-full max-w-2xl flex-col items-center px-4 py-8">
	<div
		class="neo-border neo-shadow-lg mb-8 rotate-1 px-6 py-3"
		class:bg-secondary={wizardStore.documentType === "Assignment"}
		class:bg-accent={wizardStore.documentType === "Lab Task"}
	>
		<div class="flex items-center gap-3">
			<FileText class="size-6" />
			<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
				{wizardStore.documentType} Details
			</h1>
		</div>
	</div>

	<p class="mb-8 max-w-md text-center text-base font-medium">
		Tell us about this specific {wizardStore.documentType?.toLowerCase()}.
	</p>

	<div class="w-full space-y-6">
		<!-- Document Number -->
		<div
			class="neo-border neo-shadow space-y-4 bg-card p-6"
			in:scale={{ duration: 300, delay: 100, easing: quintOut }}
		>
			<div class="mb-2 flex items-center gap-2">
				<div class="neo-border-sm bg-primary p-2">
					<Hash class="size-5" />
				</div>
				<span class="text-sm font-bold uppercase">
					{wizardStore.documentType} Number
				</span>
			</div>

			{#if wizardStore.documentType === "Assignment"}
				<AssignmentNumberSelector bind:value={formStore.document.number} />
				<p class="text-sm text-muted-foreground">Select an assignment number (1-4)</p>
			{:else}
				<Input
					id="document-number"
					name="number"
					type="number"
					placeholder="1-15"
					required
					value={formStore.document.number}
					oninput={(e) =>
						formStore.setDocumentNumber((e.target as HTMLInputElement).value)}
					class="neo-border-sm py-3 text-lg"
				/>
				<p class="text-sm text-muted-foreground">
					Which {wizardStore.documentType?.toLowerCase()} number is this?
				</p>
			{/if}
		</div>

		<!-- Total Marks -->
		<div
			class="neo-border neo-shadow space-y-4 bg-card p-6"
			in:scale={{ duration: 300, delay: 200, easing: quintOut }}
		>
			<div class="mb-2 flex items-center gap-2">
				<div class="neo-border-sm bg-secondary p-2">
					<Award class="size-5" />
				</div>
				<span class="text-sm font-bold uppercase">Total Marks</span>
			</div>

			<Input
				id="marks"
				name="marks"
				type="number"
				step="0.5"
				placeholder={wizardStore.documentType === "Assignment" ? "4" : "7.5"}
				value={formStore.document.marks}
				oninput={(e) => formStore.setDocumentMarks((e.target as HTMLInputElement).value)}
				class="neo-border-sm py-3 text-lg"
			/>
			<p class="text-sm text-muted-foreground">
				Default is {wizardStore.documentType === "Assignment" ? "4" : "7.5"} marks.
			</p>
		</div>

		<!-- Due Date -->
		<div
			class="neo-border neo-shadow space-y-4 bg-card p-6"
			in:scale={{ duration: 300, delay: 300, easing: quintOut }}
		>
			<div class="mb-2 flex items-center gap-2">
				<div class="neo-border-sm bg-accent p-2">
					<CalendarDays class="size-5" />
				</div>
				<span class="text-sm font-bold uppercase">Due Date</span>
			</div>

			<DueDate />
			<p class="text-sm text-muted-foreground">
				When is this {wizardStore.documentType?.toLowerCase()} due?
			</p>
		</div>
	</div>
</div>
