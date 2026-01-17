<script lang="ts">
    import Form from "$lib/components/form.svelte";
    import AutoForm from "$lib/components/forms/auto.svelte";
    import ImagesForm from "$lib/components/forms/images.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import SEO from "$lib/components/seo.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { Separator } from "$lib/components/ui/separator";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
</script>

<SEO
    title="lab task | zabdoc"
    description="Generate lab task PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/lab-task"
    url="https://zabdoc.xyz/lab-task"
/>

<div class="space-y-12">
    <div class="space-y-1.5">
        <h3
            class="text-2xl font-semibold leading-none tracking-tight flex items-center justify-start gap-2"
        >
            <ClipboardList />
            Lab Task
        </h3>
        <p class="text-sm text-muted-foreground">
            Fill out the information below to generate the lab task.
        </p>
    </div>

    <Form>
        <div class="grid md:grid-cols-2 gap-4">
            <AutoForm />
            <Card.Root>
                <Card.Header>
                    <Card.Title>Document</Card.Title>
                    <Card.Description
                        >Information about the lab task.</Card.Description
                    >
                </Card.Header>
                <Card.Content class="space-y-4">
                    <div class="space-y-2 hidden">
                        <Label for="document-type">Type</Label>
                        <Input name="type" value="Lab Task" />
                    </div>

                    <div class="space-y-2">
                        <Label for="marks">Total Marks</Label>
                        <Input
                            id="marks"
                            name="marks"
                            type="number"
                            step="0.5"
                            placeholder="7.5"
                            value="7.5"
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="document-number">Number</Label>
                        <Input
                            id="document-number"
                            name="number"
                            type="number"
                            placeholder="1-15"
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

        <Separator />
        <ImagesForm />
    </Form>
</div>
