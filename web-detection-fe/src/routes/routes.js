import Chart from "~/pages/chart";
import Home from "~/pages/home";

const publicRoutes = [
  { path: "/", component: Home },
  { path: "/chart", component: Chart },
];
const privateRoutes = [];

export { publicRoutes, privateRoutes };
