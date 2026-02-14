<script lang="ts">
	import { RadioGroup, RadioGroupItem } from "$lib/components/ui/radio-group/index.js";
	import { cn } from "$lib/utils.js";

	let { value = $bindable("") }: { value: string } = $props();

	const numbers = ["1", "2", "3", "4", "5"];
</script>

<RadioGroup bind:value class="grid grid-cols-5 gap-3 w-full">
	{#each numbers as num}
		<label
			class={cn(
				"neo-border neo-shadow px-4 py-6 flex items-center justify-center font-black text-lg uppercase cursor-pointer transition-all",
				value === num 
					? "bg-primary text-primary-foreground" 
					: "bg-card hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none"
			)}
		>
			<RadioGroupItem value={num} class="sr-only" />
			{num}
		</label>
	{/each}
</RadioGroup>

