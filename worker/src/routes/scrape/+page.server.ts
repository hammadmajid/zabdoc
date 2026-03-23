import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ platform }) => {
	return {
		baseURL: platform?.env.BASE_API_URL
	};
};
