<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import Button from "$lib/components/ui/button/button.svelte";
    import { browser } from "$app/environment";
    import { toast } from "svelte-sonner";
    import SEO from "$lib/components/seo.svelte";
    import { formStore } from "$lib/stores/form-store.svelte";
    import * as Select from "$lib/components/ui/select/index.js";

    let studentNameInput = $state(formStore.studentName);
    let regNoInput = $state(formStore.regNo);
    let selectedClassInput = $state(formStore.selectedClass);

    function clearLocalStorage() {
        if (!browser) return;

        try {
            // Clear specific keys used by the form store
            const keysToRemove = ["studentName", "regNo", "selectedClass"];

            keysToRemove.forEach((key) => {
                localStorage.removeItem(key);
            });

            // Reset the form store
            studentNameInput = "";
            regNoInput = "";
            selectedClassInput = "";

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
                formStore.selectedClass = selectedClassInput;
            }

            toast.success("Data updated successfully", {
                description: "Your student information has been saved.",
                position: "top-center",
            });
        } catch (error) {
            toast.error("Failed to update data", {
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
                    <select
                        id="class"
                        bind:value={selectedClassInput}
                        class="w-full px-3 py-2 neo-border-sm bg-background"
                    >
                        <option value="">Select a class</option>
                        {#each formStore.classes as classOption}
                            <option value={classOption.value}>
                                {classOption.label}
                            </option>
                        {/each}
                    </select>
                </div>

                <div class="pt-2 flex gap-2">
                    <Button onclick={updateStoredData}>
                        Save Changes
                    </Button>
                    <Button variant="destructive" onclick={clearLocalStorage}>
                        Delete Local Data
                    </Button>
                </div>
            </Card.Content>
        </Card.Root>
    </div>
</div>
