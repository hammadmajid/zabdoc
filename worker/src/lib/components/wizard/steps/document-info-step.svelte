<script lang="ts">
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import { formStore } from "$lib/stores/form-store.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import FileText from "@lucide/svelte/icons/file-text";
    import Hash from "@lucide/svelte/icons/hash";
    import Award from "@lucide/svelte/icons/award";
    import CalendarDays from "@lucide/svelte/icons/calendar-days";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-2xl mx-auto w-full">
    <div
        class="neo-border neo-shadow-lg px-6 py-3 mb-8 rotate-[1deg]"
        class:bg-secondary={wizardStore.documentType === "Assignment"}
        class:bg-accent={wizardStore.documentType === "Lab Task"}
    >
        <div class="flex items-center gap-3">
            <FileText class="size-6" />
            <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                {wizardStore.documentType} Details
            </h1>
        </div>
    </div>

    <p class="text-base font-medium max-w-md mb-8 text-center">
        Tell us about this specific {wizardStore.documentType?.toLowerCase()}.
    </p>

    <div class="w-full space-y-6">
        <!-- Document Number -->
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 100, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-primary p-2">
                    <Hash class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">
                    {wizardStore.documentType} Number
                </span>
            </div>

            <Input
                id="document-number"
                name="number"
                type="number"
                placeholder={wizardStore.documentType === "Assignment" ? "1-4" : "1-15"}
                required
                value={formStore.document.number}
                oninput={(e) => formStore.setDocumentNumber((e.target as HTMLInputElement).value)}
                class="neo-border-sm text-lg py-3"
            />
            <p class="text-sm text-muted-foreground">
                Which {wizardStore.documentType?.toLowerCase()} number is this?
            </p>
        </div>

        <!-- Total Marks -->
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 200, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-secondary p-2">
                    <Award class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">Total Marks</span>
            </div>

            <Input
                id="marks"
                name="marks"
                type="number"
                step="0.5"
                placeholder={wizardStore.documentType === "Assignment" ? "4" : "7.5"}
                value={formStore.document.marks}
                oninput={(e) => formStore.setDocumentMarks((e.target as HTMLInputElement).value)}
                class="neo-border-sm text-lg py-3"
            />
            <p class="text-sm text-muted-foreground">
                Default is {wizardStore.documentType === "Assignment" ? "4" : "7.5"} marks.
            </p>
        </div>

        <!-- Due Date -->
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 300, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-accent p-2">
                    <CalendarDays class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">Due Date</span>
            </div>

            <DueDate />
            <p class="text-sm text-muted-foreground">
                When is this {wizardStore.documentType?.toLowerCase()} due?
            </p>
        </div>
    </div>
</div>
