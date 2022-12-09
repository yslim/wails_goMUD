// import type { PageLoad } from "./$types";
import {redirect} from "@sveltejs/kit";

// export const load : PageLoad = ({params}) => {
//     console.log ('params = ', params);
// }

export function load() {
    throw redirect(307, "/home");
}
