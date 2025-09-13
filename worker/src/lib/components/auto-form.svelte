<script lang="ts">
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import { browser } from "$app/environment";
    import { data, type DataStructure } from "$lib/data/data";

    // Cookie utility functions
    function setCookie(name: string, value: string, days: number = 30) {
        if (!browser) return;
        const expires = new Date();
        expires.setTime(expires.getTime() + days * 24 * 60 * 60 * 1000);
        document.cookie = `${name}=${value};expires=${expires.toUTCString()};path=/`;
    }

    function getCookie(name: string): string {
        if (!browser) return "";
        const nameEQ = name + "=";
        const ca = document.cookie.split(";");
        for (let i = 0; i < ca.length; i++) {
            let c = ca[i];
            while (c.charAt(0) === " ") c = c.substring(1, c.length);
            if (c.indexOf(nameEQ) === 0)
                return c.substring(nameEQ.length, c.length);
        }
        return "";
    }

    let selectedClass = $state("" as keyof DataStructure | "");
    let selectedCourse = $state("");
    let studentName = $state("");
    let regNum = $state("");

    // Initialize values from cookies on component mount
    $effect(() => {
        if (browser) {
            studentName = getCookie("studentName") || "";
            regNum = getCookie("regNum") || "";
            selectedClass = (getCookie("class") as keyof DataStructure) || "";
        }
    });

    // Update cookies when values change
    $effect(() => {
        if (browser && studentName) {
            setCookie("studentName", studentName);
        }
    });

    $effect(() => {
        if (browser && regNum) {
            setCookie("regNum", regNum);
        }
    });

    $effect(() => {
        if (browser && selectedClass) {
            setCookie("class", selectedClass);
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
            This data is cached locally for your convenience and will not be
            shared.
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
            Instructor and course code are automatically filled based on your
            selection.
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
