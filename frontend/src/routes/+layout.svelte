<script>
	import '../app.css';
	import '../app.postcss';

	import { Icon } from 'flowbite-svelte-icons';

	import {
		Button,
		DarkMode,
		Dropdown,
		DropdownDivider,
		DropdownItem,
		GradientButton,
		Navbar,
		NavBrand,
		NavHamburger,
		NavLi,
		NavUl
	} from 'flowbite-svelte';

	const PreBirth = [
		{ text: 'Ultrasounds', href: '/ultrasounds' },
		{ text: 'Growing Mommy', href: '/growing' },
		{ text: 'Baby Shower / Reveal', href: '/shower_reveal' }
	];

	const Birth = [
		{ text: 'Day of Birth', href: '/dob' },
		{ text: 'Delivery (private)', href: '/delivery' }
	];

	const Stages = [
		{ text: 'Infancy (0 - 1 year)', href: 'https://gelane.quickconnect.to/mo/sharing/BaYNwj2RO' },
		{ text: 'Toddlerhood (1 - 2 years)', href: '/th' },
		{ text: 'Early Childhood (3 - 6 years)', href: '/ec' },
		{ text: 'Late Childhood (7 - 10 years)', href: '/lc' },
		{ text: 'Adolescence (11 - 19 years)', href: '/adolescence' },
		{ text: 'Early Adulthood (20 - 44 years)', href: '/ea' },
		{ text: 'Middle Adulthood (45 - 64 years)', href: '/ma' },
		{ text: 'Late Adulthood (65 years+)', href: '/la' }
	];
</script>

<svelte:head>
	<title>Gelanez Baby</title>
</svelte:head>

<div class="flex flex-col">
	<Navbar let:hidden let:toggle>
		<NavBrand href="/">
			<img src="/images/logo.png" class="mr-3 h-6 sm:h-9" alt="Docly Logo" />
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
				Gelanez Baby
			</span>
		</NavBrand>

		<div class="flex md:order-2">
			<Button color="light"
				>Account<Icon
					name="chevron-down-solid"
					class="w-3 h-3 ml-2 text-white dark:text-white"
				/></Button
			>
			<Dropdown>
				<div slot="header" class="px-4 py-2">
					<span class="block truncate text-sm font-medium">delane.jackson@gmail.com</span>
				</div>
				<DropdownItem href="/account">Settings</DropdownItem>
				<DropdownItem href="/account">Billing</DropdownItem>
				<form action="/logout" method="POST">
					<DropdownItem type="submit" slot="footer">Sign out</DropdownItem>
				</form>
			</Dropdown>
			<NavHamburger on:click={toggle} />
			<DarkMode initialTheme="light" />
		</div>
		<NavUl {hidden}>
			<NavLi href="/">
				<GradientButton outline color="purpleToPink">Home</GradientButton>
			</NavLi>
			<NavLi id="pb" class="cursor-pointer">
				<GradientButton outline color="purpleToPink"
					>Before Birth
					<Icon name="chevron-down-solid" class="w-3 h-3 ml-2 text-white dark:text-white" />
				</GradientButton>
			</NavLi>
			<NavLi id="birth" class="cursor-pointer">
				<GradientButton outline color="purpleToPink"
					>Birth
					<Icon name="chevron-down-solid" class="w-3 h-3 ml-2 text-white dark:text-white" />
				</GradientButton>
			</NavLi>
			<NavLi id="stages" class="cursor-pointer">
				<GradientButton outline color="purpleToPink"
					>Stages
					<Icon name="chevron-down-solid" class="w-3 h-3 ml-2 text-white dark:text-white" />
				</GradientButton>
			</NavLi>
			<Dropdown triggeredBy="#pb" class="w-44 z-20">
				{#each PreBirth as pb}
					<DropdownItem href={pb.href}>{pb.text}</DropdownItem>
					<DropdownDivider />
				{/each}
			</Dropdown>

			<Dropdown triggeredBy="#birth" class="w-44 z-20">
				{#each Birth as b}
					<DropdownItem href={b.href}>{b.text}</DropdownItem>
					<DropdownDivider />
				{/each}
			</Dropdown>

			<Dropdown triggeredBy="#stages" class="w-44 z-20">
				{#each Stages as stage}
					<DropdownItem href={stage.href}>{stage.text}</DropdownItem>
					<DropdownDivider />
				{/each}
			</Dropdown>
		</NavUl>
	</Navbar>
</div>

<div class="w-full flex-grow px-2 sm:px-4">
	<div class="container mx-auto">
		<slot />
	</div>
</div>
