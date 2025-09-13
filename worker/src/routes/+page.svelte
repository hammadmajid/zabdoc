<script lang="ts">
    import AutoForm from "$lib/components/forms/auto.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import ContentForm from "$lib/components/forms/content.svelte";
    import DocForm from "$lib/components/forms/doc.svelte";
    import { toast } from "svelte-sonner";

    let isLoading = $state(false);

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        isLoading = true;

        try {
            const formData = new FormData(event.target as HTMLFormElement);

            const response = await fetch("http://localhost:8080/generate", {
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
            a.download = "lab-section.pdf";
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

<div class="space-y-12 p-4 md:p-0">
    <Card.Header>
        <Card.Title>Cover</Card.Title>
        <Card.Description>
            Information that will be renderd on cover page.
        </Card.Description>
    </Card.Header>

    <form class="space-y-8" onsubmit={handleSubmit}>
        <div class="grid md:grid-cols-2 gap-4">
            <AutoForm />
            <DocForm />
        </div>

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
