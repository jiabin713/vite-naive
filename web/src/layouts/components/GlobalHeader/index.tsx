import { NLayoutHeader } from 'naive-ui';
import { defineComponent } from 'vue';
import HeaderLogo from './components/HeaderLogo';
import HeaderOperations from './components/HeaderOperations';

export default defineComponent({
  setup() {
    return () => (
      <NLayoutHeader
        bordered
        class="flex flex-row h-12 justify-between items-center"
      >
        <HeaderLogo />
        <HeaderOperations />
      </NLayoutHeader>
    );
  },
});
