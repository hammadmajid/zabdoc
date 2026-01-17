<script lang="ts">
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { formStore } from "$lib/stores/form-store.svelte";
</script>

<Card.Root>
    <Card.Header>
        <Card.Title>Images</Card.Title>
        <Card.Description>
            Upload images for your document. At least one image is required.
        </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="images">Upload Images</Label>
            <Input
                id="images"
                name="images"
                type="file"
                multiple
                required
                accept="image/jpeg,image/jpg,image/png,image/webp"
                onchange={(e) => {
                    const target = e.target as HTMLInputElement | null;
                    formStore.images = target?.files ?? null;
                }}
            />
            {#if formStore.hasImages}
                <p class="text-sm text-muted-foreground">
                    {formStore.images?.length} image{formStore.images && formStore.images.length > 1 ? "s" : ""} selected
                </p>
            {:else}
                <p class="text-sm text-destructive">
                    Please select at least one image
                </p>
            {/if}
        </div>
    </Card.Content>
</Card.Root>
