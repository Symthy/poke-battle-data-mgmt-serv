import { Controller, Get } from '@nestjs/common';
import { ApiOkResponse, ApiTags } from '@nestjs/swagger';
import { PokemonsService } from '../service/pokemons.service';
import { Pokemons } from './responseModel/pokemons';

@ApiTags('Get Poke Base Data')
@Controller('pokemons')
export class PokemonsController {
  constructor(private readonly appService: PokemonsService) { }

  @Get('')
  @ApiOkResponse({ type: [Pokemons] })
  async getPokemons(): Promise<Pokemons> {
    return this.appService.getPokemons();
  }
}
