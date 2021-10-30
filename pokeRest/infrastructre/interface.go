package repository

type repository interface {
	findAll()
	findById()
	save()
	update()
	delete()
}
