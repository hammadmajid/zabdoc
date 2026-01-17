<script lang="ts">
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import { formStore, STUDENT_LIMIT } from "$lib/stores/form-store.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import Plus from "@lucide/svelte/icons/plus";
    import X from "@lucide/svelte/icons/x";
    import User from "@lucide/svelte/icons/user";
    import Users from "@lucide/svelte/icons/users";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    // Initialize from localStorage on mount
    $effect(() => {
        formStore.initFromLocalStorage();
    });
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-2xl mx-auto w-full">
    {#if wizardStore.teamType === "individual"}
        <div class="neo-border neo-shadow-lg bg-primary px-6 py-3 mb-8 rotate-[-1deg]">
            <div class="flex items-center gap-3">
                <User class="size-6" />
                <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                    Your Information
                </h1>
            </div>
        </div>

        <p class="text-base font-medium max-w-md mb-8 text-center">
            We'll save this locally so you don't have to enter it again.
        </p>

        <div class="w-full space-y-6">
            <div
                class="neo-border neo-shadow bg-card p-6 space-y-4"
                in:scale={{ duration: 300, delay: 100, easing: quintOut }}
            >
                <div class="space-y-2">
                    <Label for="student-name" class="text-sm font-bold uppercase">Full Name</Label>
                    <Input
                        id="student-name"
                        name="studentName"
                        type="text"
                        placeholder="Enter your full name"
                        bind:value={formStore.studentName}
                        class="neo-border-sm text-lg py-3"
                    />
                </div>

                <div class="space-y-2">
                    <Label for="reg-number" class="text-sm font-bold uppercase">Registration Number</Label>
                    <Input
                        id="reg-number"
                        name="regNo"
                        type="text"
                        placeholder="Enter your registration number"
                        bind:value={formStore.regNo}
                        class="neo-border-sm text-lg py-3"
                    />
                </div>
            </div>
        </div>
    {:else}
        <div class="neo-border neo-shadow-lg bg-secondary px-6 py-3 mb-8 rotate-[1deg]">
            <div class="flex items-center gap-3">
                <Users class="size-6" />
                <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                    Group Members
                </h1>
            </div>
        </div>

        <p class="text-base font-medium max-w-md mb-8 text-center">
            Add all the members of your group (minimum 2, maximum {STUDENT_LIMIT}).
        </p>

        <div class="w-full space-y-4">
            {#each formStore.students as student, index (student.id)}
                <div
                    class="neo-border neo-shadow bg-card p-4 relative"
                    in:scale={{ duration: 300, delay: index * 50, easing: quintOut }}
                    out:scale={{ duration: 200, easing: quintOut }}
                >
                    <div class="flex items-center justify-between mb-4">
                        <div class="flex items-center gap-2">
                            <div class="neo-border-sm bg-accent px-3 py-1">
                                <span class="text-sm font-black">#{index + 1}</span>
                            </div>
                        </div>
                        <Button
                            type="button"
                            variant="outline"
                            size="sm"
                            onclick={() => formStore.removeStudent(student.id)}
                            disabled={formStore.students.length <= 2}
                            class="neo-border-sm hover:bg-destructive hover:text-white"
                        >
                            <X class="w-4 h-4" />
                        </Button>
                    </div>
                    <div class="grid sm:grid-cols-2 gap-4">
                        <div class="space-y-2">
                            <Label for="student-{student.id}-name" class="text-xs font-bold uppercase">
                                Full Name
                            </Label>
                            <Input
                                id="student-{student.id}-name"
                                type="text"
                                placeholder="Enter full name"
                                bind:value={student.name}
                                oninput={(e) => formStore.updateStudent(student.id, "name", (e.target as HTMLInputElement).value)}
                                required
                                class="neo-border-sm"
                            />
                        </div>
                        <div class="space-y-2">
                            <Label for="student-{student.id}-regNo" class="text-xs font-bold uppercase">
                                Registration Number
                            </Label>
                            <Input
                                id="student-{student.id}-regNo"
                                type="text"
                                placeholder="Enter reg number"
                                bind:value={student.regNo}
                                oninput={(e) => formStore.updateStudent(student.id, "regNo", (e.target as HTMLInputElement).value)}
                                required
                                class="neo-border-sm"
                            />
                        </div>
                    </div>
                </div>
            {/each}

            {#if formStore.students.length < STUDENT_LIMIT}
                <button
                    type="button"
                    onclick={formStore.addStudent}
                    class="w-full neo-border border-dashed bg-muted p-4 flex items-center justify-center gap-2 font-bold uppercase text-sm hover:bg-card transition-colors cursor-pointer"
                    in:scale={{ duration: 300, delay: formStore.students.length * 50, easing: quintOut }}
                >
                    <Plus class="w-5 h-5" />
                    Add Another Student
                </button>
            {/if}
        </div>
    {/if}
</div>
