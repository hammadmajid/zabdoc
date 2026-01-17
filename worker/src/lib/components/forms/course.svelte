<script lang="ts">
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import { formStore } from "$lib/stores/form-store.svelte";
</script>

<Card.Root class="">
    <Card.Header>
        <Card.Title>Course</Card.Title>
        <Card.Description>
            Instructor and course code are auto filled based on the course.
        </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="course-select">Course</Label>
            <!-- Course Selection -->
            <Select.Root
                type="single"
                name="course"
                required
                bind:value={formStore.selectedCourse}
                disabled={!formStore.selectedClass}
            >
                <Select.Trigger
                    id="course-select"
                    class="!w-full whitespace-normal min-h-[2.25rem] h-auto"
                >
                    <span class="text-left leading-tight"
                        >{formStore.courseTriggerContent}</span
                    >
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Courses</Select.Label>
                        {#each formStore.courses as course (course.value)}
                            <Select.Item
                                value={course.value}
                                label={course.label}
                            >
                                {course.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>
            {#if !formStore.selectedClass}
                <p class="text-xs text-muted-foreground">
                    Please select a class first to see available courses.
                </p>
            {/if}
        </div>

        <!-- Auto-filled fields -->
        {#if formStore.courseDetails}
            <div class="space-y-2">
                <Label for="instructor">Instructor</Label>
                <Input
                    id="instructor"
                    name="instructor"
                    type="text"
                    placeholder="Instructor"
                    value={formStore.courseDetails.instructor}
                    readonly
                    class="bg-muted"
                />
            </div>

            <div class="space-y-2">
                <Label for="course-code">Course Code</Label>
                <Input
                    id="course-code"
                    name="courseCode"
                    type="text"
                    placeholder="Course Code"
                    value={formStore.courseDetails.code}
                    readonly
                    class="bg-muted"
                />
            </div>
        {/if}
    </Card.Content>
</Card.Root>
