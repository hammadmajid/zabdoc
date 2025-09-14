import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChild<T> = T extends { child?: any } ? Omit<T, "child"> : T;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChildren<T> = T extends { children?: any } ? Omit<T, "children"> : T;
export type WithoutChildrenOrChild<T> = WithoutChildren<WithoutChild<T>>;
export type WithElementRef<T, U extends HTMLElement = HTMLElement> = T & { ref?: U | null };

export function smartName(formData: FormData): string {
	// Get form values with fallbacks
	const regNo = formData.get("regNo")?.toString() || "";
	const courseCode = formData.get("courseCode")?.toString() || "";
	const type = formData.get("type")?.toString() || "";
	const number = formData.get("number")?.toString() || "";

	// Build filename parts
	const parts: string[] = [];

	// Add registration number if available
	if (regNo) {
		parts.push(regNo);
	}

	// Add course code if available
	if (courseCode) {
		parts.push(courseCode);
	}

	// Add document type and number if available
	if (type && number) {
		// Abbreviate document types
		let typeAbbr = "";
		switch (type.toLowerCase()) {
			case "assignment":
				typeAbbr = "A";
				break;
			case "lab task":
				typeAbbr = "LT";
				break;
			case "lab project":
				typeAbbr = "LP";
				break;
			default:
				typeAbbr = type.charAt(0).toUpperCase();
		}
		parts.push(`${typeAbbr}${number}`);
	} else if (type) {
		// Just add type if number is not available
		parts.push(type.replace(/\s+/g, ""));
	} else if (number) {
		// Just add number if type is not available
		parts.push(`Doc${number}`);
	}

	// If we have parts, join them with spaces, otherwise use default
	return parts.length > 0 ? `${parts.join(" ")}.pdf` : "document.pdf";
}