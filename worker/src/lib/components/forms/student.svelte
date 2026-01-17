<script lang="ts">
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Switch } from "$lib/components/ui/switch/index.js";
    import Button from "$lib/components/ui/button/button.svelte";
    import { formStore, STUDENT_LIMIT } from "$lib/stores/form-store.svelte";
    import Plus from "@lucide/svelte/icons/plus";
    import X from "@lucide/svelte/icons/x";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    // Initialize from localStorage on mount
    $effect(() => {
        formStore.initFromLocalStorage();
    });
</script>

<Card.Root class="">
    <Card.Header>
        <div class="flex items-center justify-between">
            <div>
                <Card.Title>Student</Card.Title>
                <Card.Description
                    >This information will be cached locally.</Card.Description
                >
            </div>
            <div class="flex items-center gap-2">
                <Label for="multi-toggle" class="text-sm">Group</Label>
                <Switch id="multi-toggle" bind:checked={formStore.isMultiMode} />
            </div>
        </div>
    </Card.Header>
    <Card.Content class="space-y-4">
        {#if !formStore.isMultiMode}
            <div class="space-y-2">
                <Label for="student-name">Full Name</Label>
                <Input
                    id="student-name"
                    name="studentName"
                    type="text"
                    placeholder="Enter your full name"
                    bind:value={formStore.studentName}
                />
            </div>

            <div class="space-y-2">
                <Label for="reg-number">Registration Number</Label>
                <Input
                    id="reg-number"
                    name="regNo"
                    type="number"
                    placeholder="Enter your registration number"
                    bind:value={formStore.regNo}
                />
            </div>
        {:else}
            <div class="space-y-3">
                {#each formStore.students as student, index (student.id)}
                    <div
                        in:scale={{ duration: 400, easing: quintOut }}
                        out:scale={{ duration: 300, easing: quintOut }}
                    >
                        <Card.Root class="relative">
                            <Card.Header>
                                <div class="flex items-center justify-between">
                                    <Card.Title class="text-sm font-medium"
                                        >Student {index + 1}</Card.Title
                                    >
                                    <Button
                                        type="button"
                                        variant="outline"
                                        size="sm"
                                        onclick={() =>
                                            formStore.removeStudent(student.id)}
                                        disabled={formStore.students.length <= 2}
                                    >
                                        <X class="w-3 h-3" />
                                    </Button>
                                </div>
                            </Card.Header>
                            <Card.Content class="space-y-3">
                                <div class="space-y-2">
                                    <Label
                                        for="student-{student.id}-name"
                                        class="text-xs font-medium"
                                        >Full Name</Label
                                    >
                                    <Input
                                        id="student-{student.id}-name"
                                        type="text"
                                        placeholder="Enter full name"
                                        bind:value={student.name}
                                        oninput={(e) =>
                                            formStore.updateStudent(
                                                student.id,
                                                "name",
                                                (e.target as HTMLInputElement)
                                                    .value
                                            )}
                                        required
                                    />
                                </div>
                                <div class="space-y-2">
                                    <Label
                                        for="student-{student.id}-regNo"
                                        class="text-xs font-medium"
                                        >Registration Number</Label
                                    >
                                    <Input
                                        id="student-{student.id}-regNo"
                                        type="text"
                                        placeholder="Enter registration number"
                                        bind:value={student.regNo}
                                        oninput={(e) =>
                                            formStore.updateStudent(
                                                student.id,
                                                "regNo",
                                                (e.target as HTMLInputElement)
                                                    .value
                                            )}
                                        required
                                    />
                                </div>
                            </Card.Content>
                        </Card.Root>
                    </div>
                {/each}
                {#if formStore.students.length < STUDENT_LIMIT}
                    <div
                        in:scale={{
                            duration: 400,
                            easing: quintOut,
                            delay: 100,
                        }}
                    >
                        <Card.Root class="border-dashed">
                            <Card.Content
                                class="flex items-center justify-center py-6"
                            >
                                <Button
                                    type="button"
                                    variant="outline"
                                    class="w-full"
                                    onclick={formStore.addStudent}
                                >
                                    <Plus class="w-4 h-4 mr-2" />
                                    Add Student
                                </Button>
                            </Card.Content>
                        </Card.Root>
                    </div>
                {/if}
            </div>
        {/if}

        <div class="space-y-2">
            <Label for="class-select">Class</Label>
            <!-- Class Selection -->
            <Select.Root
                type="single"
                name="class"
                required
                bind:value={formStore.selectedClass}
            >
                <Select.Trigger
                    id="class-select"
                    class="!w-full whitespace-normal min-h-[2.25rem] h-auto"
                >
                    <span class="text-left leading-tight"
                        >{formStore.classTriggerContent}</span
                    >
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Classes</Select.Label>
                        {#each formStore.classes as classItem (classItem.value)}
                            <Select.Item
                                value={classItem.value}
                                label={classItem.label}
                            >
                                {classItem.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>
        </div>
    </Card.Content>
</Card.Root>
