import axios from 'axios';
import { NLayout } from 'naive-ui';
import { defineComponent } from 'vue';
import GlobalFooter from '../components/GlobalFooter';
import GlobalHeader from '../components/GlobalHeader';
import GlobalSider from '../components/GlobalSider';
import PageContainer from '../components/PageContainer';

export default defineComponent({
  setup() {
    return () => (
      <div class="flex flex-col h-screen w-screen">
        <GlobalHeader class="h-12 z-1 shadow-md flex-shrink-0" />
        <NLayout hasSider class="flex h-full ">
          <GlobalSider class="shadow-md flex-shrink-0 overflow-x-hidden" />
          <div class="flex-1 overflow-y-auto flex flex-col">
            <PageContainer class="flex-grow p-6" />
            <GlobalFooter />
          </div>
        </NLayout>
      </div>
    );
  },
});
