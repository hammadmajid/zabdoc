import { formStore } from "./form-store.svelte";
import { browser } from "$app/environment";

export type DocumentType = "Assignment" | "Lab Task" | null;
export type TeamType = "individual" | "group" | "blank" | null;

export type WizardStep =
	| "select-document"
	| "select-team"
	| "student-info"
	| "course-info"
	| "document-info"
	| "finalize";

function createWizardStore() {
	let currentStep = $state<WizardStep>("select-document");
	let documentType = $state<DocumentType>(null);
	let teamType = $state<TeamType>(null);
	let direction = $state<"forward" | "backward">("forward");

	// Determine the step order based on document type and team type
	const stepOrder = $derived.by(() => {
		const baseSteps: WizardStep[] = ["select-document", "select-team"];

		// Skip student-info step for blank mode
		if (teamType !== "blank") {
			baseSteps.push("student-info");
		}

		baseSteps.push("course-info", "document-info", "finalize");
		return baseSteps;
	});

	const currentStepIndex = $derived(stepOrder.indexOf(currentStep));
	const totalSteps = $derived(stepOrder.length);
	const progress = $derived(((currentStepIndex + 1) / totalSteps) * 100);
	const isFirstStep = $derived(currentStepIndex === 0);
	const isLastStep = $derived(currentStepIndex === totalSteps - 1);

	function setDocumentType(type: DocumentType) {
		documentType = type;
		if (type) {
			// Sync with form store
			formStore.setDocumentType(type);
			// Set default marks based on document type
			if (type === "Assignment") {
				formStore.setDocumentMarks("4");
			} else if (type === "Lab Task") {
				formStore.setDocumentMarks("7.5");
			}
		}
	}

	function setTeamType(type: TeamType) {
		teamType = type;
		if (type) {
			formStore.isMultiMode = type === "group";
			formStore.isBlankMode = type === "blank";
		}
	}

	function nextStep() {
		const steps = stepOrder;
		const currentIndex = steps.indexOf(currentStep);
		if (currentIndex < steps.length - 1) {
			direction = "forward";
			currentStep = steps[currentIndex + 1];
			syncToURL(true); // Use pushState for navigation
		}
	}

	function prevStep() {
		const steps = stepOrder;
		const currentIndex = steps.indexOf(currentStep);
		if (currentIndex > 0) {
			direction = "backward";
			currentStep = steps[currentIndex - 1];
			syncToURL(true); // Use pushState for navigation
		}
	}

	function goToStep(step: WizardStep) {
		const steps = stepOrder;
		const targetIndex = steps.indexOf(step);
		const currentIndex = steps.indexOf(currentStep);

		if (targetIndex !== -1) {
			direction = targetIndex > currentIndex ? "forward" : "backward";
			currentStep = step;
			syncToURL(true); // Use pushState for navigation
		}
	}

	function reset() {
		currentStep = "select-document";
		documentType = null;
		teamType = null;
		direction = "forward";
		formStore.reset();
		syncToURL(false); // Use replaceState for reset
	}

	// Sync current step to URL query parameter
	function syncToURL(usePushState = false) {
		if (!browser) return;
		const url = new URL(window.location.href);
		url.searchParams.set("step", currentStep);
		
		if (usePushState) {
			window.history.pushState({ step: currentStep }, "", url.toString());
		} else {
			window.history.replaceState({ step: currentStep }, "", url.toString());
		}
	}

	// Initialize from URL query parameter
	function initFromURL(urlStep: string | null) {
		if (!urlStep) return;

		const steps = stepOrder;
		const requestedStep = urlStep as WizardStep;

		// Check if the requested step exists in the current step order
		if (!steps.includes(requestedStep)) {
			// Invalid step, reset to first step
			currentStep = "select-document";
			syncToURL(false);
			return;
		}

		// Don't allow skipping ahead - validate that all previous steps are completeable
		const requestedIndex = steps.indexOf(requestedStep);
		for (let i = 0; i < requestedIndex; i++) {
			const step = steps[i];
			// Temporarily set to that step to check if it's valid
			const originalStep = currentStep;
			currentStep = step;
			if (!canProceed()) {
				// Can't proceed from this step, so reset to first incomplete step
				currentStep = step;
				syncToURL(false);
				return;
			}
			currentStep = originalStep;
		}

		// All validation passed, go to requested step
		currentStep = requestedStep;
	}

	// Handle browser back/forward navigation
	function handlePopState(urlStep: string | null) {
		if (!urlStep) {
			currentStep = "select-document";
			return;
		}

		const steps = stepOrder;
		const requestedStep = urlStep as WizardStep;

		// Check if the requested step exists
		if (!steps.includes(requestedStep)) {
			currentStep = "select-document";
			syncToURL(false);
			return;
		}

		const requestedIndex = steps.indexOf(requestedStep);
		const currentIndex = steps.indexOf(currentStep);

		// Set direction based on navigation
		direction = requestedIndex > currentIndex ? "forward" : "backward";

		// Allow going back to any previous step
		// But don't allow skipping ahead if requirements aren't met
		if (requestedIndex > currentIndex) {
			// Trying to go forward - validate
			for (let i = 0; i < requestedIndex; i++) {
				const step = steps[i];
				const originalStep = currentStep;
				currentStep = step;
				if (!canProceed()) {
					currentStep = step;
					syncToURL(false);
					return;
				}
				currentStep = originalStep;
			}
		}

		// All validation passed, update current step
		currentStep = requestedStep;
	}

	function canProceed(): boolean {
		switch (currentStep) {
			case "select-document":
				return documentType !== null;
			case "select-team":
				return teamType !== null;
			case "student-info":
				if (teamType === "blank") {
					return true; // Always allow proceeding for blank mode
				} else if (teamType === "individual") {
					return formStore.studentName.trim() !== "" && formStore.regNo.trim() !== "";
				} else {
					return formStore.students.every(
						(s) => s.name.trim() !== "" && s.regNo.trim() !== ""
					);
				}
			case "course-info":
				return formStore.selectedClass !== "" && formStore.selectedCourse !== "";
			case "document-info":
				return formStore.document.number !== "" && formStore.document.date !== "";
			case "finalize":
				return true;
			default:
				return false;
		}
	}

	return {
		get currentStep() {
			return currentStep;
		},
		get documentType() {
			return documentType;
		},
		get teamType() {
			return teamType;
		},
		get direction() {
			return direction;
		},
		get currentStepIndex() {
			return currentStepIndex;
		},
		get totalSteps() {
			return totalSteps;
		},
		get progress() {
			return progress;
		},
		get isFirstStep() {
			return isFirstStep;
		},
		get isLastStep() {
			return isLastStep;
		},
		get stepOrder() {
			return stepOrder;
		},

		setDocumentType,
		setTeamType,
		nextStep,
		prevStep,
		goToStep,
		reset,
		canProceed,
		syncToURL,
		initFromURL,
		handlePopState
	};
}

export const wizardStore = createWizardStore();
