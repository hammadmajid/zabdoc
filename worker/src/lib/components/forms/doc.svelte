<script lang="ts">
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import * as Select from "$lib/components/ui/select/index.js";

    const types = ["Assignment", "Lab Task"];

    let value = $state(types[0]);
    let marks = $state("");

    const triggerContent = $derived(
        types.find((f) => f === value) ?? "Select type",
    );

    // Auto-set marks based on document type
    $effect(() => {
        if (value === "Assignment") {
            marks = "4";
        } else if (value === "Lab Task" || value === "Lab Project") {
            marks = "7.5";
        }
    });
</script>

<Card.Root>
    <Card.Header>
        <Card.Title>Document</Card.Title>
        <Card.Description>Information about the document.</Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="document-type">Type</Label>
            <Select.Root type="single" name="type" bind:value>
                <Select.Trigger class="w-full">{triggerContent}</Select.Trigger>
                <Select.Content>
                    {#each types as type (type)}
                        <Select.Item value={type} label={type}>
                            {type}
                        </Select.Item>
                    {/each}
                </Select.Content>
            </Select.Root>
        </div>

        <div class="space-y-2 hidden">
            <Label for="marks">Marks</Label>
            <Input
                id="marks"
                name="marks"
                type="number"
                step="0.5"
                placeholder="4"
                bind:value={marks}
            />
        </div>

        <div class="space-y-2">
            <Label for="document-number">Number</Label>
            <Input
                id="document-number"
                name="number"
                type="number"
                placeholder="1"
                required
            />
        </div>

        <div class="space-y-2">
            <Label for="due-date">Due Date</Label>
            <DueDate />
        </div>
    </Card.Content>
</Card.Root>
