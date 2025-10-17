<script lang="ts">
    import { questionsStore, scalingStore, type SubmissionResult } from "../lib/stores";

    let {submissions,name,question}: {submissions: SubmissionResult[],name:string,question:number} = $props();
    let q = $questionsStore[question]
    let back_total = submissions.filter(x => !x.front)
    let back_correct = back_total.filter(x => x.answer in q.correct)
    let front_total = submissions.filter(x => x.front)
    let front_correct = front_total.filter(x => x.answer in q.correct)

    let back_ratio = back_correct.length / back_total.length
    let front_ratio = front_correct.length / front_total.length

</script>

<div class="flex:column">
    <span>{name}</span>
    <span>Anzahl: {submissions.length}</span>
    <span>Hinten Anzahl: {back_total.length}</span>
    <span>Vorne Anzahl: {front_total.length}</span>
    <span>Hinten Richtig: {back_correct.length}</span>
    <span>Vorne Richtig: {front_correct.length}</span>
    <span>Hinten Verhältnis: {back_ratio}</span>
    <span>Vorne Verhältnis: {front_ratio}</span>
    <span>Punkte: {front_ratio * back_ratio * $scalingStore}</span>
</div>