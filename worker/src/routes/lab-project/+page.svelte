<script lang="ts">
    import AutoForm from "$lib/components/forms/auto.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import ContentForm from "$lib/components/forms/content.svelte";
    import { toast } from "svelte-sonner";
    import { smartName } from "$lib/utils";
    import { Label } from "$lib/components/ui/label";
    import { Input } from "$lib/components/ui/input";
    import * as Select from "$lib/components/ui/select/index";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import Microscope from "@lucide/svelte/icons/microscope";
    import { Separator } from "$lib/components/ui/separator";

    let isLoading = $state(false);
    let selectedProjectType = $state("");

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        isLoading = true;

        try {
            const formData = new FormData(event.target as HTMLFormElement);

            // Shitty hack to check if we are in localhost or prod
            const apiUrl = window.location.host.includes("localhost")
                ? "http://localhost:8080/generate"
                : "https://api.zabdoc.xyz/generate";
            const response = await fetch(apiUrl, {
                method: "POST",
                body: formData,
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // Handle PDF download
            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.style.display = "none";
            a.href = url;
            a.download = smartName(formData);
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url);
            document.body.removeChild(a);
        } catch (error) {
            toast.error("Failed to generate PDF", {
                description: `${error}`,
                position: "top-center",
            });
        }
        isLoading = false;
    }
</script>

<svelte:head>
    <title>lab project | zabdoc</title>
</svelte:head>

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

    <form class="space-y-8" onsubmit={handleSubmit}>
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
        <ContentForm />

        <Card.Root class="">
            <Card.Header>
                <Card.Description>
                    By clicking the button below you accept the <a
                        href="/about"
                        class="text-primary underline hover:no-underline"
                        >Terms and Privacy Policy</a
                    >.
                </Card.Description>
            </Card.Header>
            <Card.CardContent class="space-y-4">
                <Button type="submit" class="w-full" disabled={isLoading}>
                    {#if isLoading}
                        <Loader2 class="mr-2 h-4 w-4 animate-spin" />
                        Generating Document...
                    {:else}
                        Generate Document
                    {/if}
                </Button>
            </Card.CardContent>
        </Card.Root>
    </form>
</div>
