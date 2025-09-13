<script lang="ts">
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import { browser } from "$app/environment";
    import { data, type DataStructure } from "$lib/data/data";
    import { toast } from "svelte-sonner";

    // LocalStorage utility functions
    function setLocalStorage(key: string, value: string) {
        if (!browser) return;
        try {
            localStorage.setItem(key, value);
        } catch (error) {
            toast.warning("Failed to save to localStorage", {
                description: error as string,
                position: "top-center",
            });
        }
    }

    function getLocalStorage(key: string): string {
        if (!browser) return "";
        try {
            return localStorage.getItem(key) || "";
        } catch (error) {
            toast.warning("Failed to read from localStorage", {
                description: error as string,
                position: "top-center",
            });
            return "";
        }
    }

    let selectedClass = $state("" as keyof DataStructure | "");
    let selectedCourse = $state("");
    let studentName = $state("");
    let regNum = $state("");

    // Initialize values from localStorage on component mount
    $effect(() => {
        if (browser) {
            studentName = getLocalStorage("studentName") || "";
            regNum = getLocalStorage("regNum") || "";
            selectedClass =
                (getLocalStorage("class") as keyof DataStructure) || "";
        }
    });

    // Update localStorage when values change
    $effect(() => {
        if (browser && studentName) {
            setLocalStorage("studentName", studentName);
        }
    });

    $effect(() => {
        if (browser && regNum) {
            setLocalStorage("regNum", regNum);
        }
    });

    $effect(() => {
        if (browser && selectedClass) {
            setLocalStorage("class", selectedClass);
        }
    });

    const classes = $derived(
        Object.keys(data).map((className) => ({
            value: className,
            label: className,
        })),
    );

    const courses = $derived(
        selectedClass && selectedClass in data
            ? Object.keys(data[selectedClass]).map((courseName) => ({
                  value: courseName,
                  label: courseName,
              }))
            : [],
    );

    const courseDetails = $derived(
        selectedClass &&
            selectedCourse &&
            selectedClass in data &&
            selectedCourse in data[selectedClass]
            ? data[selectedClass][selectedCourse]
            : null,
    );

    const classTriggerContent = $derived(
        classes.find((c) => c.value === selectedClass)?.label ?? "Select class",
    );

    const courseTriggerContent = $derived(
        courses.find((c) => c.value === selectedCourse)?.label ??
            "Select course",
    );

    // Reset course selection when class changes
    $effect(() => {
        if (selectedClass) {
            selectedCourse = "";
        }
    });
</script>

<Card.Root class="">
    <Card.Header>
        <Card.Title>Student</Card.Title>
        <Card.Description>
            This data is stored locally for your convenience.
        </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="student-name">Full Name</Label>
            <Input
                id="student-name"
                name="studentName"
                type="text"
                placeholder="Enter your full name"
                bind:value={studentName}
                required
            />
        </div>

        <div class="space-y-2">
            <Label for="reg-number">Registration Number</Label>
            <Input
                id="reg-number"
                name="regNo"
                type="number"
                placeholder="Enter your registration number"
                bind:value={regNum}
                required
            />
        </div>

        <div class="space-y-2">
            <Label for="class-select">Class</Label>
            <!-- Class Selection -->
            <Select.Root
                type="single"
                name="class"
                required
                bind:value={selectedClass}
            >
                <Select.Trigger id="class-select" class="w-full">
                    {classTriggerContent}
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Classes</Select.Label>
                        {#each classes as classItem (classItem.value)}
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

<Card.Root class="">
    <Card.Header>
        <Card.Title>Course</Card.Title>
        <Card.Description>
            Instructor and course code are auto filled based on the
            course.
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
                bind:value={selectedCourse}
                disabled={!selectedClass}
            >
                <Select.Trigger id="course-select" class="w-full">
                    {courseTriggerContent}
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Courses</Select.Label>
                        {#each courses as course (course.value)}
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
            {#if !selectedClass}
                <p class="text-xs text-muted-foreground">
                    Please select a class first to see available courses.
                </p>
            {/if}
        </div>

        <!-- Auto-filled fields -->
        {#if courseDetails}
            <div class="space-y-2">
                <Label for="instructor">Instructor</Label>
                <Input
                    id="instructor"
                    name="instructor"
                    type="text"
                    placeholder="Instructor"
                    value={courseDetails.instructor}
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
                    value={courseDetails.code}
                    readonly
                    class="bg-muted"
                />
            </div>
        {/if}
    </Card.Content>
</Card.Root>
