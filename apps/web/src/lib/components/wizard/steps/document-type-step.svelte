<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import FileText from "@lucide/svelte/icons/file-text";
	import ClipboardList from "@lucide/svelte/icons/clipboard-list";
</script>

<div class="flex flex-col items-center justify-center px-4 py-8 text-center">
	<div class="neo-border neo-shadow-lg mb-8 -rotate-1 bg-primary px-8 py-4">
		<h1 class="text-3xl font-black tracking-tight uppercase md:text-4xl">
			What do you need?
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
			class="neo-border neo-shadow flex cursor-pointer flex-col items-center gap-4 bg-card px-10 py-8 text-xl font-black tracking-wide uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:bg-secondary hover:shadow-none"
		>
			<div class="neo-border-sm bg-secondary p-4">
				<FileText class="size-10" />
			</div>
			<div class="flex flex-col items-center gap-1">
				<span>Assignment</span>
				<span class="text-xs font-medium text-muted-foreground normal-case">Select document type</span>
			</div>
		</button>
		<button
			type="button"
			disabled
			class="neo-border neo-shadow flex cursor-not-allowed flex-col items-center gap-4 bg-card px-10 py-8 text-xl font-black tracking-wide uppercase opacity-50 transition-all"
		>
			<div class="neo-border-sm bg-accent p-4">
				<ClipboardList class="size-10" />
			</div>
			<div class="flex flex-col items-center gap-1">
				<span>Lab Task</span>
				<span class="text-xs font-medium text-muted-foreground normal-case">Under maintenance</span>
			</div>
		</button>
	</div>
</div>
