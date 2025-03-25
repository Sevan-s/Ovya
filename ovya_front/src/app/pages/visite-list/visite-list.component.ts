import { Component, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { combineLatest } from 'rxjs';
import { startWith } from 'rxjs/operators';

interface Visite {
  id: number;
  date_start: string;
  date_end: string;
  acq_id: number;
  ccial_id: number;
  dossier_id: number;
  canceled: boolean;
  acq_name?: string;
}

interface Acq {
  ID: number;
  Nom: string;
}

@Component({
  selector: 'visite-list',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './visite-list.component.html',
  styleUrl: './visite-list.component.css'
})
export class VisiteListComponent {
  http = inject(HttpClient);

  visites: Visite[] = [];
  filteredVisites: Visite[] = [];
  allAcq: Acq[] = [];
  dataIsEmpty: boolean= false;
  searchControl = new FormControl('');
  filterTypeControl = new FormControl('acq');

  constructor() {
    this.loadData();

    combineLatest([
      this.searchControl.valueChanges.pipe(startWith('acq')),
      this.filterTypeControl.valueChanges.pipe(startWith(this.filterTypeControl.value))
    ]).subscribe(([term, type]) => {
      const value = (term || '').toLowerCase();

      this.filteredVisites = this.visites.filter(v => {
        const dateStart = new Date(v.date_start).toLocaleDateString('fr-FR');
        const dateEnd = new Date(v.date_end).toLocaleDateString('fr-FR');

        switch (type) {
          case 'acq':
            return v.acq_name?.toLowerCase().includes(value);
          case 'date':
            return dateStart.includes(value) || dateEnd.includes(value);
          case 'ccial':
            return v.ccial_id.toString().includes(value);
          case 'dossier':
            return v.dossier_id.toString().includes(value);
          default:
            return true;
        }
      });
    });
  }

  loadData() {
    this.http.get<Visite[]>('http://localhost:8080/api/v1/visite/all').subscribe(visites => {
      this.visites = visites;

      console.log("visite", this.visites)
      if (this.visites !== null) {
        this.loadAcqs();
      this.dataIsEmpty = false;
    }
      else
        this.dataIsEmpty = true
    });

  }

  loadAcqs() {
    this.http.get<Acq[]>('http://localhost:8080/api/v1/acq/all').subscribe(acqs => {
      this.allAcq = acqs;

      this.visites = this.visites.map(v => {
        const match = this.allAcq.find(a => a.ID === v.acq_id);
        return {
          ...v,
          acq_name: match ? match.Nom : 'Inconnu'
        };
      });

      this.filteredVisites = [...this.visites];
    });
  }
}