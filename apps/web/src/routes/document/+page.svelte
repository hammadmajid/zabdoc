<script lang="ts">
	import { onMount } from "svelte";
	import Wizard from "$lib/components/wizard/wizard.svelte";
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import { fade } from "svelte/transition";
	import ArrowLeft from "@lucide/svelte/icons/arrow-left";
	import type { PageProps } from "./$types";
	let { data }: PageProps = $props();

	// Reset wizard when landing on document page
	onMount(() => {
		wizardStore.reset();
	});
</script>

<div class="flex min-h-[70vh] flex-col">
	<!-- Back to Home Button - Always visible -->
	<div class="mb-4" in:fade={{ duration: 300 }}>
		<a
			href="/"
			class="neo-border neo-shadow flex w-fit items-center gap-2 bg-card px-4 py-2 text-sm font-bold uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
		>
			<ArrowLeft class="size-4" />
			Back to Home
		</a>
	</div>

	{#if wizardStore.currentStep === "select-document"}
		<!-- Hero Section - shown only on first step -->
		<div class="mb-4 text-center" in:fade={{ duration: 300 }}>
			<div class="neo-border neo-shadow-lg mb-4 inline-block -rotate-2 bg-primary px-8 py-4">
				<h1 class="text-4xl font-black tracking-tight uppercase md:text-6xl">
					Generate Document
				</h1>
			</div>
			<p class="mx-auto max-w-xl text-lg font-medium">
				Create assignment cover sheets and lab tasks in PDF format.
			</p>
		</div>
	{/if}

	<!-- Wizard Component -->
	<div class="flex-1">
		<Wizard baseURL={data.baseURL} />
	</div>
</div>
