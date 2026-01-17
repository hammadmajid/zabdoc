<script lang="ts">
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { formStore } from "$lib/stores/form-store.svelte";
    import { smartName } from "$lib/utils";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import { toast } from "svelte-sonner";

    let isLoading = $state(false);

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        isLoading = true;

        try {
            // Build FormData from the store
            const formData = formStore.buildFormData();

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

    let { children } = $props();
</script>

<form class="space-y-8" onsubmit={handleSubmit}>
    {@render children?.()}

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
