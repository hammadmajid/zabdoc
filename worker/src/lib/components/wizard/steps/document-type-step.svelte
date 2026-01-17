<script lang="ts">
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import FileText from "@lucide/svelte/icons/file-text";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
</script>

<div class="flex flex-col items-center justify-center text-center px-4 py-8">
    <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 mb-8 rotate-[-2deg]">
        <h1 class="text-4xl md:text-5xl font-black uppercase tracking-tight">
            What do you need?
        </h1>
    </div>

    <p class="text-lg font-medium max-w-md mb-8">
        Select the type of document you want to generate.
    </p>

    <div class="flex flex-col sm:flex-row gap-6">
        <button
            type="button"
            onclick={() => {
                wizardStore.setDocumentType("Assignment");
                wizardStore.nextStep();
            }}
            class="neo-border neo-shadow bg-secondary text-secondary-foreground px-10 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer"
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
            class="neo-border neo-shadow bg-accent text-accent-foreground px-10 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer"
        >
            <ClipboardList class="size-8" />
            Lab Task
        </button>
    </div>
</div>
