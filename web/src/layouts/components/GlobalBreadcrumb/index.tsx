import { NBreadcrumb } from 'naive-ui';
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const routeStore = useRouteStore();

export default defineComponent({
  setup() {
    return () => <NBreadcrumb>breadcrumbs.map</NBreadcrumb>;
  },
});
