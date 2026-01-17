<script lang="ts">
    import { formStore } from "$lib/stores/form-store.svelte";
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import BookOpen from "@lucide/svelte/icons/book-open";
    import GraduationCap from "@lucide/svelte/icons/graduation-cap";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    // Initialize from localStorage on mount
    $effect(() => {
        formStore.initFromLocalStorage();
    });
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-2xl mx-auto w-full">
    <div class="neo-border neo-shadow-lg bg-secondary px-6 py-3 mb-8 rotate-[-1deg]">
        <div class="flex items-center gap-3">
            <BookOpen class="size-6" />
            <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                Course Details
            </h1>
        </div>
    </div>

    <p class="text-base font-medium max-w-md mb-8 text-center">
        Select your class and course. Instructor details will be filled automatically.
    </p>

    <div class="w-full space-y-6">
        <!-- Class Selection -->
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 100, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-primary p-2">
                    <GraduationCap class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">Select Class</span>
            </div>

            <Select.Root
                type="single"
                name="class"
                required
                bind:value={formStore.selectedClass}
            >
                <Select.Trigger class="!w-full neo-border-sm whitespace-normal min-h-[3rem] h-auto text-lg">
                    <span class="text-left leading-tight">{formStore.classTriggerContent}</span>
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Classes</Select.Label>
                        {#each formStore.classes as classItem (classItem.value)}
                            <Select.Item value={classItem.value} label={classItem.label}>
                                {classItem.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>
        </div>

        <!-- Course Selection -->
        <div
            class="neo-border neo-shadow bg-card p-6 space-y-4"
            in:scale={{ duration: 300, delay: 200, easing: quintOut }}
        >
            <div class="flex items-center gap-2 mb-2">
                <div class="neo-border-sm bg-accent p-2">
                    <BookOpen class="size-5" />
                </div>
                <span class="font-bold uppercase text-sm">Select Course</span>
            </div>

            <Select.Root
                type="single"
                name="course"
                required
                bind:value={formStore.selectedCourse}
                disabled={!formStore.selectedClass}
            >
                <Select.Trigger class="!w-full neo-border-sm whitespace-normal min-h-[3rem] h-auto text-lg">
                    <span class="text-left leading-tight">{formStore.courseTriggerContent}</span>
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Courses</Select.Label>
                        {#each formStore.courses as course (course.value)}
                            <Select.Item value={course.value} label={course.label}>
                                {course.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>

            {#if !formStore.selectedClass}
                <p class="text-sm text-muted-foreground bg-muted neo-border-sm px-3 py-2">
                    Please select a class first to see available courses.
                </p>
            {/if}
        </div>

        <!-- Auto-filled Course Details -->
        {#if formStore.courseDetails}
            <div
                class="neo-border neo-shadow bg-muted p-6 space-y-4"
                in:scale={{ duration: 300, easing: quintOut }}
            >
                <div class="flex items-center gap-2 mb-2">
                    <div class="neo-border-sm bg-secondary p-2">
                        <ClipboardList class="size-5" />
                    </div>
                    <span class="font-bold uppercase text-sm">Auto-filled Details</span>
                </div>

                <div class="grid sm:grid-cols-2 gap-4">
                    <div class="space-y-2">
                        <Label for="instructor" class="text-xs font-bold uppercase">Instructor</Label>
                        <Input
                            id="instructor"
                            name="instructor"
                            type="text"
                            value={formStore.courseDetails.instructor}
                            readonly
                            class="neo-border-sm bg-card"
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="course-code" class="text-xs font-bold uppercase">Course Code</Label>
                        <Input
                            id="course-code"
                            name="courseCode"
                            type="text"
                            value={formStore.courseDetails.code}
                            readonly
                            class="neo-border-sm bg-card"
                        />
                    </div>
                </div>
            </div>
        {/if}
    </div>
</div>
