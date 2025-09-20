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
            // Clear specific keys used by the auto.svelte
            const keysToRemove = ["studentName", "regNum", "class"];

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

<div class="space-y-6 max-w-2xl mx-auto">
    <div>
        <h1 class="text-3xl font-bold tracking-tight">Settings</h1>
        <p class="text-muted-foreground">Manage your preferences and data.</p>
    </div>

    <!-- Theme Settings -->
    <Card.Root>
        <Card.Header>
            <Card.Title>Appearance</Card.Title>
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
            <Card.Title>Storage</Card.Title>
            <Card.Description>
                Manage your locally stored data.
            </Card.Description>
        </Card.Header>
        <Card.Content class="space-y-4">
            <div class="space-y-2">
                <h4 class="text-sm font-medium">Local Data</h4>
                <p class="text-sm text-muted-foreground">
                    Your name, registration number, and class selection are
                    stored locally in your browser for convenience.
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
