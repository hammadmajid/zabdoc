<script lang="ts">
    import { formStore, type ImageItem } from "$lib/stores/form-store.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import ImageIcon from "@lucide/svelte/icons/image";
    import Upload from "@lucide/svelte/icons/upload";
    import X from "@lucide/svelte/icons/x";
    import GripVertical from "@lucide/svelte/icons/grip-vertical";
    import Plus from "@lucide/svelte/icons/plus";
    import Trash2 from "@lucide/svelte/icons/trash-2";
    import { scale, fade } from "svelte/transition";
    import { flip } from "svelte/animate";
    import { quintOut } from "svelte/easing";

    let fileInput: HTMLInputElement;
    let dragOverZone = $state(false);
    let draggingId = $state<string | null>(null);
    let dragOverId = $state<string | null>(null);

    function handleFileSelect(e: Event) {
        const target = e.target as HTMLInputElement;
        if (target.files && target.files.length > 0) {
            formStore.addImages(target.files);
            target.value = "";
        }
    }

    function handleDrop(e: DragEvent) {
        e.preventDefault();
        dragOverZone = false;

        if (e.dataTransfer?.files && e.dataTransfer.files.length > 0) {
            const validFiles = Array.from(e.dataTransfer.files).filter((file) =>
                ["image/jpeg", "image/jpg", "image/png", "image/webp"].includes(file.type)
            );
            if (validFiles.length > 0) {
                formStore.addImages(validFiles);
            }
        }
    }

    function handleDragOver(e: DragEvent) {
        e.preventDefault();
        dragOverZone = true;
    }

    function handleDragLeave(e: DragEvent) {
        e.preventDefault();
        dragOverZone = false;
    }

    // Reorder drag handlers
    function handleItemDragStart(e: DragEvent, id: string) {
        draggingId = id;
        if (e.dataTransfer) {
            e.dataTransfer.effectAllowed = "move";
            e.dataTransfer.setData("text/plain", id);
        }
    }

    function handleItemDragEnd() {
        draggingId = null;
        dragOverId = null;
    }

    function handleItemDragOver(e: DragEvent, id: string) {
        e.preventDefault();
        if (draggingId && draggingId !== id) {
            dragOverId = id;
        }
    }

    function handleItemDragLeave() {
        dragOverId = null;
    }

    function handleItemDrop(e: DragEvent, targetId: string) {
        e.preventDefault();
        if (draggingId && draggingId !== targetId) {
            const fromIndex = formStore.imageItems.findIndex((img) => img.id === draggingId);
            const toIndex = formStore.imageItems.findIndex((img) => img.id === targetId);
            if (fromIndex !== -1 && toIndex !== -1) {
                formStore.reorderImages(fromIndex, toIndex);
            }
        }
        draggingId = null;
        dragOverId = null;
    }

    function triggerFileInput() {
        fileInput?.click();
    }
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-3xl mx-auto w-full">
    <div class="neo-border neo-shadow-lg bg-accent px-6 py-3 mb-8 rotate-[-1deg]">
        <div class="flex items-center gap-3">
            <ImageIcon class="size-6" />
            <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                Upload Images
            </h1>
        </div>
    </div>

    <p class="text-base font-medium max-w-md mb-8 text-center">
        Add screenshots or photos for your document. Drag to reorder them.
    </p>

    <!-- Hidden file input -->
    <input
        bind:this={fileInput}
        type="file"
        multiple
        accept="image/jpeg,image/jpg,image/png,image/webp"
        onchange={handleFileSelect}
        class="hidden"
    />

    <div class="w-full space-y-6">
        <!-- Drop Zone -->
        <div
            class="neo-border neo-shadow bg-card p-6 transition-all duration-200"
            class:bg-primary={dragOverZone}
            class:scale-[1.02]={dragOverZone}
            in:scale={{ duration: 300, delay: 100, easing: quintOut }}
            role="button"
            tabindex="0"
            ondrop={handleDrop}
            ondragover={handleDragOver}
            ondragleave={handleDragLeave}
            onclick={triggerFileInput}
            onkeydown={(e) => e.key === "Enter" && triggerFileInput()}
        >
            <div class="flex flex-col items-center justify-center py-8 gap-4 cursor-pointer">
                <div
                    class="neo-border-sm p-4 transition-transform duration-200"
                    class:bg-card={dragOverZone}
                    class:bg-primary={!dragOverZone}
                    class:rotate-[-5deg]={dragOverZone}
                >
                    <Upload class="size-10" />
                </div>
                <div class="text-center">
                    <p class="font-bold uppercase text-lg">
                        {dragOverZone ? "Drop images here!" : "Click or drag images here"}
                    </p>
                    <p class="text-sm text-muted-foreground mt-1">
                        JPEG, PNG, WebP supported
                    </p>
                </div>
            </div>
        </div>

        <!-- Image Grid -->
        {#if formStore.hasImages}
            <div
                class="neo-border neo-shadow bg-card p-4"
                in:scale={{ duration: 300, easing: quintOut }}
            >
                <div class="flex items-center justify-between mb-4">
                    <div class="flex items-center gap-2">
                        <div class="neo-border-sm bg-secondary px-3 py-1">
                            <span class="text-sm font-black uppercase">
                                {formStore.imageItems.length} image{formStore.imageItems.length > 1 ? "s" : ""}
                            </span>
                        </div>
                        <span class="text-sm text-muted-foreground">Drag to reorder</span>
                    </div>
                    <Button
                        variant="outline"
                        size="sm"
                        onclick={formStore.clearImages}
                        class="neo-border-sm hover:bg-destructive hover:text-white"
                    >
                        <Trash2 class="size-4 mr-2" />
                        Clear All
                    </Button>
                </div>

                <div class="grid grid-cols-2 sm:grid-cols-3 gap-4">
                    {#each formStore.imageItems as item, index (item.id)}
                        <div
                            class="relative group"
                            animate:flip={{ duration: 300, easing: quintOut }}
                            draggable="true"
                            role="listitem"
                            ondragstart={(e) => handleItemDragStart(e, item.id)}
                            ondragend={handleItemDragEnd}
                            ondragover={(e) => handleItemDragOver(e, item.id)}
                            ondragleave={handleItemDragLeave}
                            ondrop={(e) => handleItemDrop(e, item.id)}
                        >
                            <div
                                class="neo-border-sm overflow-hidden bg-muted aspect-square transition-all duration-200 cursor-grab active:cursor-grabbing"
                                class:opacity-50={draggingId === item.id}
                                class:ring-4={dragOverId === item.id}
                                class:ring-primary={dragOverId === item.id}
                                class:scale-105={dragOverId === item.id}
                            >
                                <img
                                    src={item.previewUrl}
                                    alt="Upload {index + 1}"
                                    class="w-full h-full object-cover"
                                />
                            </div>

                            <!-- Order badge -->
                            <div
                                class="absolute -top-2 -left-2 neo-border-sm bg-primary px-2 py-1 z-10"
                            >
                                <span class="text-xs font-black">#{index + 1}</span>
                            </div>

                            <!-- Drag handle -->
                            <div
                                class="absolute top-2 right-2 neo-border-sm bg-card/90 p-1.5 opacity-0 group-hover:opacity-100 transition-opacity cursor-grab"
                            >
                                <GripVertical class="size-4" />
                            </div>

                            <!-- Remove button -->
                            <button
                                type="button"
                                onclick={() => formStore.removeImage(item.id)}
                                class="absolute -bottom-2 -right-2 neo-border-sm bg-destructive text-white p-1.5 opacity-0 group-hover:opacity-100 transition-opacity hover:scale-110 z-10"
                            >
                                <X class="size-4" />
                            </button>
                        </div>
                    {/each}

                    <!-- Add more button -->
                    <button
                        type="button"
                        onclick={triggerFileInput}
                        class="neo-border-sm border-dashed aspect-square flex flex-col items-center justify-center gap-2 bg-muted hover:bg-secondary transition-colors cursor-pointer"
                    >
                        <Plus class="size-8 text-muted-foreground" />
                        <span class="text-xs font-bold uppercase text-muted-foreground">Add More</span>
                    </button>
                </div>
            </div>
        {:else}
            <div
                class="neo-border border-dashed bg-muted p-8"
                in:scale={{ duration: 300, easing: quintOut }}
            >
                <div class="flex flex-col items-center justify-center gap-3 text-muted-foreground">
                    <ImageIcon class="size-12" />
                    <p class="font-bold uppercase">No images yet</p>
                    <p class="text-sm text-center">
                        Upload screenshots or photos to include in your document
                    </p>
                </div>
            </div>
        {/if}
    </div>
</div>
