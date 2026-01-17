<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import Button from "$lib/components/ui/button/button.svelte";
    import Theme from "$lib/components/theme.svelte";
    import { browser } from "$app/environment";
    import { toast } from "svelte-sonner";
    import SEO from "$lib/components/seo.svelte";

    function clearLocalStorage() {
        if (!browser) return;

        try {
            // Clear specific keys used by the form store
            const keysToRemove = ["studentName", "regNo"];

            keysToRemove.forEach((key) => {
                localStorage.removeItem(key);
            });

            toast.success("Local data cleared successfully", {
                description: "Your saved student information has been removed.",
                position: "top-center",
            });
        } catch (error) {
            toast.error("Failed to clear local data", {
                description: `${error}`,
                position: "top-center",
            });
        }
    }
</script>

<SEO
    title="settings | zabdoc"
    description="Manage your zabdoc preferences and local data."
    canonical="https://zabdoc.xyz/settings"
    url="https://zabdoc.xyz/settings"
/>

<div class="space-y-6 max-w-5xl mx-auto">
    <div>
        <h1 class="text-3xl font-black uppercase tracking-tight">Settings</h1>
        <p class="font-medium">Manage your preferences and data.</p>
    </div>

    <div class="w-full grid md:grid-cols-2 gap-4">
        <!-- Theme Settings -->
        <Card.Root>
            <Card.Header>
                <Card.Title class="font-black uppercase">Appearance</Card.Title>
                <Card.Description>
                    Customize how zabdoc looks on your device.
                </Card.Description>
            </Card.Header>
            <Card.Content>
                <Theme />
            </Card.Content>
        </Card.Root>

        <!-- Storage Settings -->
        <Card.Root>
            <Card.Header>
                <Card.Title class="font-black uppercase">Storage</Card.Title>
                <Card.Description>
                    Manage your locally stored data.
                </Card.Description>
            </Card.Header>
            <Card.Content class="space-y-4">
                <div class="space-y-2">
                    <h4 class="text-sm font-bold">Local Data</h4>
                    <p class="text-sm">
                        Your name and registration number are stored locally in
                        your browser for convenience.
                    </p>
                </div>

                <div class="pt-2">
                    <Button variant="destructive" onclick={clearLocalStorage}>
                        Delete Local Data
                    </Button>
                </div>
            </Card.Content>
        </Card.Root>
    </div>
</div>
