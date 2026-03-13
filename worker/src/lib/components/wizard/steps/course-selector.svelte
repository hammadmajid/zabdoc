<script lang="ts">
	import { RadioGroup, RadioGroupItem } from "$lib/components/ui/radio-group/index.js";
	import { cn } from "$lib/utils.js";

	interface Course {
		value: string;
		label: string;
	}

	let { value = $bindable(""), courses = [], disabled = false }: { value: string; courses: Course[]; disabled: boolean } = $props();
</script>

<RadioGroup bind:value {disabled} class="flex flex-col gap-3 w-full">
	{#each courses as course (course.value)}
		<label
			class={cn(
				"neo-border neo-shadow px-6 py-4 flex items-center gap-3 font-bold uppercase cursor-pointer transition-all",
				disabled && "opacity-50 cursor-not-allowed",
				value === course.value
					? "bg-accent text-accent-foreground"
					: "bg-card hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none"
			)}
		>
			<RadioGroupItem value={course.value} class="sr-only" {disabled} />
			{course.label}
		</label>
	{/each}
</RadioGroup>
