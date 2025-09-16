<script lang="ts">
    import AutoForm from "$lib/components/forms/auto.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import DocForm from "$lib/components/forms/doc.svelte";
    import { toast } from "svelte-sonner";
    import { smartName } from "$lib/utils";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Select from "$lib/components/ui/select/index.js";

    let isLoading = $state(false);

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
                description: error as string,
                position: "top-center",
            });
        } finally {
            isLoading = false;
        }
    }
</script>

<svelte:head>
    <title>zabdoc</title>
</svelte:head>

<div class="space-y-12">
    <div class="space-y-1.5">
        <h3 class="text-2xl font-semibold leading-none tracking-tight">
            Document
        </h3>
        <p class="text-sm text-muted-foreground">
            Fill out the information below to generate the document.
        </p>
    </div>

    <form class="space-y-8" onsubmit={handleSubmit}>
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
                        <Input  name="type" value="Assignment" />
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
