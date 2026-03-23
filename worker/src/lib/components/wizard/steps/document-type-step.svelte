<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import FileText from "@lucide/svelte/icons/file-text";
	import ClipboardList from "@lucide/svelte/icons/clipboard-list";

	// Check if it's between 12 AM and 3 AM
	const isLateNight = $derived.by(() => {
		const now = new Date();
		const hours = now.getHours();

		// Check if time is between 00:00 and 02:59 (12 AM to 3 AM)
		return hours >= 0 && hours < 3;
	});
</script>

<div class="flex flex-col items-center justify-center px-4 py-8 text-center">
	<div class="neo-border neo-shadow-lg mb-8 -rotate-2 bg-primary px-8 py-4">
		<h1 class="text-lg font-black tracking-tight uppercase md:text-2xl">
			{#if isLateNight}
				Your sleep schedule is cooked.
			{:else}
				What do you need?
			{/if}
		</h1>
	</div>

	<p class="mb-8 max-w-md text-lg font-medium">
		Select the type of document you want to generate.
	</p>

	<div class="flex flex-col gap-6 sm:flex-row">
		<button
			type="button"
			onclick={() => {
				wizardStore.setDocumentType("Assignment");
				wizardStore.nextStep();
			}}
			class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-secondary px-10 py-6 text-xl font-black tracking-wide text-secondary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
		>
			<FileText class="size-8" />
			Assignment
		</button>
		<button
			type="button"
			onclick={() => {
				wizardStore.setDocumentType("Lab Task");
				wizardStore.nextStep();
			}}
			class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-accent px-10 py-6 text-xl font-black tracking-wide text-accent-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
		>
			<ClipboardList class="size-8" />
			Lab Task
		</button>
	</div>
</div>
