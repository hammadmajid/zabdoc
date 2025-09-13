<script lang="ts">
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import * as Select from "$lib/components/ui/select/index.js";

    const types = [
        { value: "assignment", label: "Assignment" },
        { value: "lab-task", label: "Lab Task" },
        { value: "lab-project", label: "Lab Project" },
    ];

    let value = $state(types[0].value);

    const triggerContent = $derived(
        types.find((f) => f.value === value)?.label ?? "Select type",
    );
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
                <Select.Trigger class="w-full"
                    >{triggerContent}</Select.Trigger
                >
                <Select.Content>
                    {#each types as type (type.value)}
                        <Select.Item
                            value={type.value}
                            label={type.label}
                            disabled={type.value === "grapes"}
                        >
                            {type.label}
                        </Select.Item>
                    {/each}
                </Select.Content>
            </Select.Root>
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
