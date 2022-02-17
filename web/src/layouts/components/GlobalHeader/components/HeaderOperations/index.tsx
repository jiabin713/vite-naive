import { NButton, NIcon, NSpace, NTooltip } from 'naive-ui';
import { defineComponent } from 'vue';
import AvatarDropDown from '../AvatarDropDown';
import { Email, Moon, Search } from '@vicons/carbon';

export default defineComponent({
  setup() {
    return () => (
      <div class="mr-4">
        <NSpace wrap={false} size={'large'} class="justify-end items-center">
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <Search />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>搜索</span>,
            }}
          />
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <Email />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>搜索</span>,
            }}
          />
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <Moon />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>深色模式</span>,
            }}
          />
          <AvatarDropDown />
        </NSpace>
      </div>
    );
  },
});
