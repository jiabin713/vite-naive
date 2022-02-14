import {
  NH2,
  NLayout,
  NLayoutContent,
  NLayoutFooter,
  NLayoutHeader,
  NLayoutSider,
} from 'naive-ui';
import { defineComponent } from 'vue';
import GlobalFooter from '../components/GlobalFooter';
import GlobalHeader from '../components/GlobalHeader';
import GlobalSider from '../components/GlobalSider';
import PageContainer from '../components/PageContainer';

export default defineComponent({
  setup() {
    return () => (
      <div class="h-screen w-screen flex flex-col ">
        <div class="h-12 bg-red-500">
          <NLayoutHeader bordered>header</NLayoutHeader>
        </div>
        <div class="flex flex-row flew-nowrap h-full box-border">
          <NLayout hasSider>
            <NLayoutSider bordered>
              <NH2>sd</NH2>
              <NH2>sd</NH2>
              <NH2>sd</NH2>
              <NH2>sd</NH2>
              <NH2>sd</NH2>
            </NLayoutSider>
            <NLayout class="flex flex-col">
              <NLayoutContent>content</NLayoutContent>
              <NLayoutFooter>sdfd</NLayoutFooter>
            </NLayout>
          </NLayout>
        </div>
      </div>
    );
  },
});
