import { TypeCompatibility, TypeCompatibilityTable, Types } from '@/models/typeCompatibility';
import { Injectable } from '@nestjs/common';

@Injectable()
export class TypesService {
  getTypes(): Types {
    return new Types();
  }

  getTypeCompatibilityTable(): TypeCompatibilityTable {
    return new TypeCompatibilityTable();
  }

  getTypeCompatibility(type: string): TypeCompatibility {
    return new TypeCompatibility();
  }
}
