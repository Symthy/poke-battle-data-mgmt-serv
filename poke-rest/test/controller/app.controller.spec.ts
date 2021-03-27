import { Test, TestingModule } from '@nestjs/testing';
import { PokemonsController } from '../../src/controller/pokemons.controller';
import { PokemonsService } from '../../src/service/pokemons.service';

describe('AppController', () => {
  let appController: PokemonsController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [PokemonsController],
      providers: [PokemonsService],
    }).compile();

    appController = app.get<PokemonsController>(PokemonsController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(appController.getPokemons()).toBe('Hello World!');
    });
  });
});
