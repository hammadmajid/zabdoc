<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import SEO from "$lib/components/seo.svelte";

    onMount(() => {
        // Reset wizard and set document type, then redirect
        wizardStore.reset();
        wizardStore.setDocumentType("Assignment");
        wizardStore.nextStep();
        goto("/", { replaceState: true });
    });
</script>

<SEO
    title="assignment | zabdoc"
    description="Generate assignment cover sheets for SZABIST students."
    canonical="https://zabdoc.xyz/"
    url="https://zabdoc.xyz/"
/>

<div class="flex flex-col items-center justify-center min-h-[50vh] text-center">
    <div class="neo-border neo-shadow bg-primary px-6 py-3 animate-pulse">
        <p class="text-xl font-black uppercase">Redirecting...</p>
    </div>
</div>
