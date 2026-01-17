<script lang="ts">
    import { formStore } from "$lib/stores/form-store.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import ImageIcon from "@lucide/svelte/icons/image";
    import Upload from "@lucide/svelte/icons/upload";
    import Check from "@lucide/svelte/icons/check";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-2xl mx-auto w-full">
    <div class="neo-border neo-shadow-lg bg-accent px-6 py-3 mb-8 rotate-[-1deg]">
        <div class="flex items-center gap-3">
            <ImageIcon class="size-6" />
            <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                Upload Images
            </h1>
        </div>
    </div>

    <p class="text-base font-medium max-w-md mb-8 text-center">
        Add screenshots or photos for your lab task. These will be included in the PDF.
    </p>

    <div class="w-full space-y-6">
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 100, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-primary p-2">
                    <Upload class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">Select Images</span>
            </div>

            <div class="relative">
                <Input
                    id="images"
                    name="images"
                    type="file"
                    multiple
                    required
                    accept="image/jpeg,image/jpg,image/png,image/webp"
                    onchange={(e) => {
                        const target = e.target as HTMLInputElement | null;
                        formStore.images = target?.files ?? null;
                    }}
                    class="neo-border-sm text-lg py-3 file:mr-4 file:py-2 file:px-4 file:neo-border-sm file:bg-primary file:text-primary-foreground file:font-bold file:uppercase file:text-sm file:cursor-pointer"
                />
            </div>

            <p class="text-sm text-muted-foreground">
                Accepted formats: JPEG, PNG, WebP
            </p>
        </div>

        {#if formStore.hasImages}
            <div
                class="neo-border neo-shadow bg-secondary p-6"
                in:scale={{ duration: 300, easing: quintOut }}
            >
                <div class="flex items-center gap-3">
                    <div class="neo-border-sm bg-card p-2">
                        <Check class="size-5 text-green-600" />
                    </div>
                    <div>
                        <p class="font-bold uppercase text-sm">
                            {formStore.images?.length} image{formStore.images && formStore.images.length > 1 ? "s" : ""} selected
                        </p>
                        <p class="text-sm text-muted-foreground">
                            Ready to be included in your lab task PDF.
                        </p>
                    </div>
                </div>
            </div>
        {:else}
            <div
                class="neo-border border-dashed bg-muted p-6"
                in:scale={{ duration: 300, easing: quintOut }}
            >
                <div class="flex items-center justify-center gap-3 text-muted-foreground">
                    <ImageIcon class="size-8" />
                    <p class="font-medium">No images selected yet</p>
                </div>
            </div>
        {/if}
    </div>
</div>
