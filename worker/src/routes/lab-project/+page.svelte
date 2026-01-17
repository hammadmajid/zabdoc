<script lang="ts">
    import Form from "$lib/components/form.svelte";
    import AutoForm from "$lib/components/forms/auto.svelte";
    import ImagesForm from "$lib/components/forms/images.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import SEO from "$lib/components/seo.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import * as Select from "$lib/components/ui/select/index";
    import { Separator } from "$lib/components/ui/separator";
    import Microscope from "@lucide/svelte/icons/microscope";

    let selectedProjectType = $state("");
</script>

<SEO
    title="lab project | zabdoc"
    description="Generate lab project PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/lab-project"
    url="https://zabdoc.xyz/lab-project"
/>

<div class="space-y-12">
    <div class="space-y-1.5">
        <h3
            class="text-2xl font-semibold leading-none tracking-tight flex items-center justify-start gap-2"
        >
            <Microscope />
            Lab Project
        </h3>
        <p class="text-sm text-muted-foreground">
            Fill out the information below to generate the lab project.
        </p>
    </div>

    <Form>
        <div class="grid md:grid-cols-2 gap-4">
            <AutoForm />
            <Card.Root>
                <Card.Header>
                    <Card.Title>Project</Card.Title>
                    <Card.Description
                        >Information about the lab project.</Card.Description
                    >
                </Card.Header>
                <Card.Content class="space-y-4">
                    <div class="space-y-2">
                        <Label for="project-type">Project Type</Label>
                        <Select.Root
                            type="single"
                            name="type"
                            required
                            bind:value={selectedProjectType}
                        >
                            <Select.Trigger id="project-type" class="w-full">
                                <span
                                    >{selectedProjectType ||
                                        "Select Project Type"}</span
                                >
                            </Select.Trigger>
                            <Select.Content>
                                <Select.Group>
                                    <Select.Label>Project Types</Select.Label>
                                    <Select.Item
                                        value="Mid Term Project"
                                        label="Mid Term Project"
                                    >
                                        Mid Term Project
                                    </Select.Item>
                                    <Select.Item
                                        value="Final Term Project"
                                        label="Final Term Project"
                                    >
                                        Final Term Project
                                    </Select.Item>
                                </Select.Group>
                            </Select.Content>
                        </Select.Root>
                    </div>

                    <div class="space-y-2">
                        <Label for="project-title">Project Title</Label>
                        <Input
                            id="project-title"
                            name="projectTitle"
                            type="text"
                            placeholder="Enter your project title"
                            required
                        />
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
