<script lang="ts">
    import { Checkbox } from "$lib/components/ui/checkbox/index.js";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import * as Select from "$lib/components/ui/select/index.js";

    const fonts = [
        // value is the name of file stored in R2 bucket
        { value: "times.ttf", label: "Times New Roman" },
        { value: "atkinson.tff", label: " Atkinson Hyperlegible" },
    ];

    let value = $state(fonts[0].value);

    const triggerContent = $derived(
        fonts.find((f) => f.value === value)?.label ?? "Select font",
    );
</script>

<Card.Root>
    <Card.Header>
        <Card.Title>Options</Card.Title>
        <Card.Description>Control the generated output.</Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="document-font">Font</Label>
            <Select.Root type="single" name="font" bind:value>
                <Select.Trigger class="w-full">{triggerContent}</Select.Trigger>
                <Select.Content>
                    {#each fonts as font (font)}
                        <Select.Item value={font.value} label={font.label}>
                            {font.label}
                        </Select.Item>
                    {/each}
                </Select.Content>
            </Select.Root>
        </div>

        <div class="flex items-center gap-3">
            <Checkbox id="syntax" disabled />
            <Label for="syntax">Syntax highlighting</Label>
        </div>
    </Card.Content>
</Card.Root>
