import { Controller, DefaultValuePipe, Get, Param, Query } from '@nestjs/common';
import { ApiOkResponse, ApiQuery, ApiTags } from '@nestjs/swagger';
import { Move, Moves } from '../models/moves';
import { MovesService } from '../provider/moves.service';

@ApiTags('Get Poke Base Data')
@Controller('moves')
export class MovesController {
  constructor(private readonly moveService: MovesService) { }

  @Get('')
  @ApiQuery({
    name: 'name',
    description: "move name",
    required: false,
    type: String
  })
  @ApiQuery({
    name: 'type',
    description: "move type",
    required: false,
    type: String
  })
  @ApiOkResponse({ type: [Moves] })
  async getMoves(
    @Query('name', new DefaultValuePipe('')) name: string,
    @Query('type', new DefaultValuePipe('')) type: string
  ): Promise<Moves> {
    return this.moveService.getMoves();
  }

  @Get(':id')
  @ApiOkResponse({ type: [Move] })
  async getMoveById(@Param('id') id: string): Promise<Move> {
    return this.moveService.getMove(id);
  }
}
