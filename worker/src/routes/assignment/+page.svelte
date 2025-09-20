<script lang="ts">
    import Form from "$lib/components/form.svelte";
    import AutoForm from "$lib/components/forms/auto.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import SEO from "$lib/components/seo.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import FileText from "@lucide/svelte/icons/file-text";
</script>

<SEO
    title="assignment | zabdoc"
    description="Generate assignment cover sheets for SZABIST students."
    canonical="https://zabdoc.xyz/assignment"
    url="https://zabdoc.xyz/assignment"
/>

<div class="space-y-12">
    <div class="space-y-1.5">
        <h3
            class="text-2xl font-semibold leading-none tracking-tight flex items-center justify-start gap-2"
        >
            <FileText />
            Assignment
        </h3>
        <p class="text-sm text-muted-foreground">
            Fill out the information below to generate the assignment cover.
        </p>
    </div>

    <Form>
        <div class="grid md:grid-cols-2 gap-4">
            <AutoForm />

            <Card.Root>
                <Card.Header>
                    <Card.Title>Document</Card.Title>
                    <Card.Description
                        >Information about the assignment.</Card.Description
                    >
                </Card.Header>
                <Card.Content class="space-y-4">
                    <div class="space-y-2 hidden">
                        <Label for="document-type">Type</Label>
                        <Input name="type" value="Assignment" />
                    </div>

                    <div class="space-y-2">
                        <Label for="marks">Total Marks</Label>
                        <Input
                            id="marks"
                            name="marks"
                            type="number"
                            step="0.5"
                            placeholder="4"
                            value="4"
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="document-number">Number</Label>
                        <Input
                            id="document-number"
                            name="number"
                            type="number"
                            placeholder="1-4"
                            required
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="due-date">Due Date</Label>
                        <DueDate />
                    </div>
                </Card.Content>
            </Card.Root>
        </div>
    </Form>
</div>
