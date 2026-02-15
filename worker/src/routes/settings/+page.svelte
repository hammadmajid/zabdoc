<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import Button from "$lib/components/ui/button/button.svelte";
    import { browser } from "$app/environment";
    import SEO from "$lib/components/seo.svelte";
    import { formStore } from "$lib/stores/form-store.svelte";
    import * as Select from "$lib/components/ui/select/index.js";
    import { onMount } from "svelte";
    import CheckCircle2 from "@lucide/svelte/icons/check-circle-2";

    let studentNameInput = $state("");
    let regNoInput = $state("");
    let selectedClassInput = $state<string>("");
    let saveButtonState = $state<"idle" | "saved">("idle");
    let deleteButtonState = $state<"idle" | "deleted">("idle");

    // Initialize from localStorage on mount
    onMount(() => {
        formStore.initFromLocalStorage();
        studentNameInput = formStore.studentName;
        regNoInput = formStore.regNo;
        selectedClassInput = formStore.selectedClass;
    });

    function clearLocalStorage() {
        if (!browser) return;

        try {
            // Clear specific keys used by the form store
            const keysToRemove = ["studentName", "regNo", "selectedClass"];

            keysToRemove.forEach((key) => {
                localStorage.removeItem(key);
            });

            // Reset the form store
            formStore.studentName = "";
            formStore.regNo = "";
            formStore.selectedClass = "";

            // Reset input values
            studentNameInput = "";
            regNoInput = "";
            selectedClassInput = "";

            // Show deleted state
            deleteButtonState = "deleted";
            
            // Reset after 2 seconds
            setTimeout(() => {
                deleteButtonState = "idle";
            }, 2000);
        } catch (error) {
            console.error("Failed to clear local data:", error);
        }
    }

    function updateStoredData() {
        if (!browser) return;

        try {
            // Update student name and registration number
            if (studentNameInput.trim()) {
                formStore.studentName = studentNameInput;
            }
            if (regNoInput.trim()) {
                formStore.regNo = regNoInput;
            }
            if (selectedClassInput) {
                formStore.selectedClass = selectedClassInput as any;
            }

            // Show saved state
            saveButtonState = "saved";
            
            // Reset after 2 seconds
            setTimeout(() => {
                saveButtonState = "idle";
            }, 2000);
        } catch (error) {
            console.error("Failed to update data:", error);
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
                    <label for="name" class="text-sm font-bold">Student Name</label>
                    <input
                        id="name"
                        type="text"
                        bind:value={studentNameInput}
                        placeholder="Enter your name"
                        class="w-full px-3 py-2 neo-border-sm bg-background"
                    />
                </div>

                <div class="space-y-2">
                    <label for="regno" class="text-sm font-bold">Registration Number</label>
                    <input
                        id="regno"
                        type="text"
                        bind:value={regNoInput}
                        placeholder="Enter your registration number"
                        class="w-full px-3 py-2 neo-border-sm bg-background"
                    />
                </div>

                <div class="space-y-2">
                    <label for="class" class="text-sm font-bold">Class</label>
                    <Select.Root
                        type="single"
                        bind:value={selectedClassInput}
                    >
                        <Select.Trigger id="class" class="w-full">
                            <span>
                                {formStore.classes.find((c) => c.value === selectedClassInput)?.label ||
                                    "Select a class"}
                            </span>
                        </Select.Trigger>
                        <Select.Content>
                            <Select.Group>
                                <Select.Label>Available Classes</Select.Label>
                                {#each formStore.classes as classOption (classOption.value)}
                                    <Select.Item value={classOption.value} label={classOption.label}>
                                        {classOption.label}
                                    </Select.Item>
                                {/each}
                            </Select.Group>
                        </Select.Content>
                    </Select.Root>
                </div>

                <div class="pt-2 flex gap-2">
                    <Button 
                        onclick={() => updateStoredData()}
                        disabled={saveButtonState === "saved"}
                    >
                        {#if saveButtonState === "saved"}
                            <CheckCircle2 class="size-4 mr-2" />
                            Saved
                        {:else}
                            Save Changes
                        {/if}
                    </Button>
                    <Button 
                        variant="destructive" 
                        onclick={() => clearLocalStorage()}
                        disabled={deleteButtonState === "deleted"}
                    >
                        {#if deleteButtonState === "deleted"}
                            <CheckCircle2 class="size-4 mr-2" />
                            Deleted
                        {:else}
                            Delete Local Data
                        {/if}
                    </Button>
                </div>
            </Card.Content>
        </Card.Root>
    </div>
</div>
