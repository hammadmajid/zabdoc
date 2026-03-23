<script lang="ts">
	import { RadioGroup, RadioGroupItem } from "$lib/components/ui/radio-group/index.js";
	import { cn } from "$lib/utils.js";

	interface Course {
		value: string;
		label: string;
	}

	let {
		value = $bindable(""),
		courses = [],
		disabled = false
	}: { value: string; courses: Course[]; disabled: boolean } = $props();
</script>

<RadioGroup bind:value {disabled} class="flex w-full flex-col gap-3">
	{#each courses as course (course.value)}
		<label
			class={cn(
				"neo-border neo-shadow flex cursor-pointer items-center gap-3 px-6 py-4 font-bold uppercase transition-all",
				disabled && "cursor-not-allowed opacity-50",
				value === course.value
					? "bg-accent text-accent-foreground"
					: "bg-card hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
			)}
		>
			<RadioGroupItem value={course.value} class="sr-only" {disabled} />
			{course.label}
		</label>
	{/each}
</RadioGroup>
