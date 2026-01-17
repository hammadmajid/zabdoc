<script lang="ts">
    import SEO from "$lib/components/seo.svelte";
    import Wizard from "$lib/components/wizard/wizard.svelte";
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import { fade } from "svelte/transition";

    // Reset wizard when landing on homepage
    $effect(() => {
        // Only reset on fresh page load, not during navigation within wizard
        return () => {
            // Cleanup if needed when leaving page
        };
    });
</script>

<SEO
    title="zabdoc"
    description="Generate assignment and lab task PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/"
    url="https://zabdoc.xyz/"
/>

<div class="min-h-[70vh] flex flex-col">
    {#if wizardStore.currentStep === "select-document"}
        <!-- Hero Section - shown only on first step -->
        <div class="text-center mb-4" in:fade={{ duration: 300 }}>
            <div class="sr-only neo-border neo-shadow-lg bg-primary px-8 py-4 inline-block rotate-[-2deg] mb-4">
                <h1 class="text-4xl md:text-6xl font-black uppercase tracking-tight">
                    zabdoc
                </h1>
            </div>
            <p class="text-lg font-medium max-w-xl mx-auto">
                web application for
                <span class="bg-secondary px-2 py-0.5 neo-border-sm font-bold">SZABIST*</span>
                students to generate assignment and lab tasks in PDF format.
            </p>
        </div>
    {/if}

    <!-- Wizard Component -->
    <div class="flex-1">
        <Wizard />
    </div>

    <!-- Disclaimer - shown only on first step -->
    {#if wizardStore.currentStep === "select-document"}
        <div class="text-center mt-4" in:fade={{ duration: 300, delay: 200 }}>
            <p class="text-sm font-medium bg-muted neo-border-sm px-4 py-2 max-w-md inline-block">
                *this project is <em>not sponsored, affiliated, endorsed, or approved by</em> SZABIST.
                <a href="/about#disclaimer" class="underline font-bold hover:text-primary">Learn more.</a>
            </p>
        </div>
    {/if}
</div>
