<script lang="ts">
    import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { Textarea } from "$lib/components/ui/textarea";
    import Plus from "@lucide/svelte/icons/plus";
    import X from "@lucide/svelte/icons/x";
    import Separator from "../ui/separator/separator.svelte";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    interface Section {
        id: number;
        content: string;
        files: FileList | null;
    }

    let sections: Section[] = $state([{ id: 0, content: "", files: null }]);
    let sectionCounter = $state(1);

    function addsection() {
        sectionCounter++;
        sections = [
            ...sections,
            { id: sectionCounter, content: "", files: null },
        ];
    }

    function removesection(sectionId: number) {
        if (sections.length > 1) {
            sections = sections.filter((section) => section.id !== sectionId);
        }
    }

    function updatesectionContent(sectionId: number, content: string) {
        sections = sections.map((section) =>
            section.id === sectionId ? { ...section, content } : section,
        );
    }

    function updatesectionFiles(sectionId: number, files: FileList | null) {
        sections = sections.map((section) =>
            section.id === sectionId ? { ...section, files } : section,
        );
    }
</script>

<div class="grid md:grid-cols-2 gap-4">
    {#each sections as section, index (section.id)}
        <div
            in:scale={{ duration: 400, easing: quintOut }}
            out:scale={{ duration: 300, easing: quintOut }}
        >
            <Card.Root class="relative">
                <Card.Header>
                    <div class="flex items-center justify-between">
                        <Card.Title>Section {index + 1}</Card.Title>
                        {#if sections.length > 1}
                            <Button
                                type="button"
                                variant="outline"
                                size="sm"
                                onclick={() => removesection(section.id)}
                            >
                                <X />
                            </Button>
                        {/if}
                    </div>
                    <Card.Description
                        >Add section content and attach supporting files (images
                        only).</Card.Description
                    >
                </Card.Header>
                <Card.Content class="space-y-4">
                    <div class="space-y-2">
                        <Label for="section-{index}-content"
                            >Section Content</Label
                        >
                        <Textarea
                            id="section-{index}-content"
                            rows={3}
                            name="section-{index}-content"
                            placeholder="Enter your content here. [Markdown formatting is supported]"
                            bind:value={section.content}
                            oninput={(e) => {
                                const target =
                                    e.target as HTMLTextAreaElement | null;
                                if (target)
                                    updatesectionContent(
                                        section.id,
                                        target.value,
                                    );
                            }}
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="section-{index}-files">Attach Files</Label>
                        <Input
                            id="section-{index}-files"
                            name="section-{index}-files"
                            type="file"
                            multiple
                            accept="image/jpeg,image/jpg,image/png,image/webp"
                            onchange={(e) => {
                                const target =
                                    e.target as HTMLInputElement | null;
                                updatesectionFiles(
                                    section.id,
                                    target?.files ?? null,
                                );
                            }}
                        />
                    </div>
                </Card.Content>
            </Card.Root>
        </div>
    {/each}

    {#if sections.length < 15}
        <div in:scale={{ duration: 400, easing: quintOut, delay: 100 }}>
            <Card.Root class="md:min-h-[250px] h-full border-dashed">
                <Card.Content class="h-full flex items-center justify-center">
                    <Button
                        class="w-full h-full min-h-[200px] flex flex-col gap-2"
                        type="button"
                        variant="outline"
                        onclick={addsection}
                    >
                        <Plus />
                        <span>Add New section</span>
                        <span class="text-sm text-muted-foreground"
                            >Click to add another section</span
                        >
                    </Button>
                </Card.Content>
            </Card.Root>
        </div>
    {/if}
</div>
