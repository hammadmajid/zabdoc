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
		
		// Initialize from URL if step param exists
		const urlParams = new URLSearchParams(window.location.search);
		const urlStep = urlParams.get("step");
		if (urlStep) {
			wizardStore.initFromURL(urlStep);
		} else {
			// Sync initial step to URL
			wizardStore.syncToURL();
		}

		// Listen to browser back/forward button
		const handlePopState = () => {
			const urlParams = new URLSearchParams(window.location.search);
			const urlStep = urlParams.get("step");
			wizardStore.handlePopState(urlStep);
		};

		window.addEventListener("popstate", handlePopState);

		return () => {
			window.removeEventListener("popstate", handlePopState);
		};
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

	<!-- Wizard Component -->
	<div class="flex-1">
		<Wizard baseURL={data.baseURL} />
	</div>
</div>
