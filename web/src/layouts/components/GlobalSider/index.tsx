import { h, ref } from 'vue';
import {
  CloudServiceManagement,
  MeterAlt,
  TwoFactorAuthentication,
} from '@vicons/carbon';
import { NIcon, NLayoutSider, NMenu } from 'naive-ui';
import { MenuMixedOption } from 'naive-ui/lib/menu/src/interface';
import { defineComponent } from 'vue';

function renderIcon(icon: any) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions: MenuMixedOption[] = [
  {
    label: 'Dashboard',
    key: 'dashboard',
    icon: renderIcon(MeterAlt),
    children: [
      {
        label: '分析页',
        key: 'analysis',
      },
      {
        label: '监控页',
        key: 'monitor',
      },
      {
        label: '工作台',
        key: 'workplace',
      },
    ],
  },
  {
    label: '权限管理',
    key: 'permission-settings',
    icon: renderIcon(TwoFactorAuthentication),
    children: [
      {
        label: '用户管理',
        key: 'staff',
      },
      {
        label: '角色管理',
        key: 'role',
      },
      {
        label: '组织管理',
        key: 'organization',
      },
      {
        label: '职位管理',
        key: 'job',
      },
      {
        label: '菜单管理',
        key: 'menu',
      },
    ],
  },
  {
    label: '系统管理',
    key: 'system-settings',
    icon: renderIcon(CloudServiceManagement),
    children: [
      { label: '字典管理', key: 'dictionary' },
      { label: '日志管理', key: 'log' },
    ],
  },
];

export default defineComponent({
  setup() {
    const collapsed = ref(false);

    return () => (
      <NLayoutSider
        collapseMode="width"
        width={208}
        collapsedWidth={64}
        bordered
        showTrigger
        nativeScrollbar={false}
        collapsed={collapsed.value}
        onCollapse={() => (collapsed.value = true)}
        onExpand={() => (collapsed.value = false)}
      >
        <NMenu
          collapsed={collapsed.value}
          collapsedWidth={64}
          collapsedIconSize={22}
          indent={18}
          options={menuOptions}
        />
      </NLayoutSider>
    );
  },
});
