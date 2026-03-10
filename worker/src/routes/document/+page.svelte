<script lang="ts">
    import { onMount } from "svelte";
    import Wizard from "$lib/components/wizard/wizard.svelte";
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import { fade } from "svelte/transition";
    import ArrowLeft from "@lucide/svelte/icons/arrow-left";

    // Reset wizard when landing on document page
    onMount(() => {
        wizardStore.reset();
    });
</script>

<div class="min-h-[70vh] flex flex-col">
    <!-- Back to Home Button - Always visible -->
    <div class="mb-4" in:fade={{ duration: 300 }}>
        <a
            href="/"
            class="neo-border neo-shadow bg-card px-4 py-2 font-bold uppercase text-sm flex items-center gap-2 hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all inline-flex w-fit"
        >
            <ArrowLeft class="size-4" />
            Back to Home
        </a>
    </div>

    {#if wizardStore.currentStep === "select-document"}
        <!-- Hero Section - shown only on first step -->
        <div class="text-center mb-4" in:fade={{ duration: 300 }}>
            <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 inline-block rotate-[-2deg] mb-4">
                <h1 class="text-4xl md:text-6xl font-black uppercase tracking-tight">
                    Generate Document
                </h1>
            </div>
            <p class="text-lg font-medium max-w-xl mx-auto">
                Create assignment cover sheets and lab tasks in PDF format.
            </p>
        </div>
    {/if}

    <!-- Wizard Component -->
    <div class="flex-1">
        <Wizard />
    </div>
</div>
