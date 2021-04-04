import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

async function main() {

  await prisma.user.create({
    data: {
      name: 'SYM',
      email: 'sym@prisma.io',
      profile: {
        create: { bio: 'I like pokemon' },
      },
    },
  })

  const allUsers = await prisma.user.findMany()

  console.log(allUsers)
}

main()
  .catch((e) => {
    throw e
  })
  .finally(async () => {
    await prisma.$disconnect()
  })
