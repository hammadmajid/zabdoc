<script lang="ts">
	import { RadioGroup, RadioGroupItem } from "$lib/components/ui/radio-group/index.js";
	import { cn } from "$lib/utils.js";

	let { value = $bindable("") }: { value: string } = $props();

	const numbers = ["1", "2", "3", "4"];
</script>

<RadioGroup bind:value class="grid w-full grid-cols-4 gap-3">
	{#each numbers as num}
		<label
			class={cn(
				"neo-border neo-shadow flex cursor-pointer items-center justify-center px-4 py-6 text-lg font-black uppercase transition-all",
				value === num
					? "bg-primary text-primary-foreground"
					: "bg-card hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
			)}
		>
			<RadioGroupItem value={num} class="sr-only" />
			{num}
		</label>
	{/each}
</RadioGroup>
